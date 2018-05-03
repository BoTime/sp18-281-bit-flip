package handler

import (
	"cmpe281/common/output"
	"net/http"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
	"cmpe281/inventory/model"
	"log"
	"cmpe281/common"
	"encoding/json"
	"strconv"
	"time"
)

func (ctx *RequestContext) CreateAllocation(w http.ResponseWriter, r *http.Request) {
	userId, err := gocql.ParseUUID(common.GetUserId(r))
	if err != nil {
		output.WriteErrorMessage(w, http.StatusUnauthorized, "User Session Not Valid")
		return
	}

	storeId, err := ctx.getStoreIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
		return
	}

	dbShard := SelectShard(storeId, ctx.Database)

	var allocationRequest model.CreateAllocationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&allocationRequest); err != nil {
		log.Println(err)
		output.WriteErrorMessage(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if len(allocationRequest.Products) == 0 {
		log.Println("No Products Provided")
		output.WriteErrorMessage(w, http.StatusBadRequest, "No Products Provided")
		return
	}

	// -- Start: Parse ID from Item Name
	{
		defaultUUID := gocql.UUID{}
		for _, product := range allocationRequest.Products {
			if product.Id.String() == defaultUUID.String() {
				// Cassandra doesn't support IN clause for non-primary searches :(
				// We have to look up each name individually
				query, names := qb.Select("products").Columns("id").Where(qb.Eq("name")).Limit(1).ToCql()
				q := gocqlx.Query(dbShard.Query(query), names).BindStruct(model.ProductDetails{
					Name: product.Name,
				})
				var productId model.ProductDetails
				if err := gocqlx.Get(&productId, q.Query); err != nil {
					log.Println("Failed to lookup Product IDs: ", err)
					output.WriteErrorMessage(w, http.StatusInternalServerError, "Failed to lookup Product IDs")
					return
				}

				product.Id = productId.Id
			}
		}
	}
	// -- End: Parse ID from Item Name

	log.Println(allocationRequest)

	// -- Start: Retrieve Current Inventory from DB --
	productIds := make([]gocql.UUID, len(allocationRequest.Products))
	for i := 0; i < len(allocationRequest.Products); i++ {
		productIds[i] = allocationRequest.Products[i].Id
	}

	query, names := qb.Select("inventory").Where(qb.Eq("store_id"), qb.In("id")).ToCql()
	q := gocqlx.Query(dbShard.Query(query), names).BindMap(qb.M{
		"store_id": storeId,
		"id": productIds,
	})

	inventory := make([]model.InventoryDetails, 0)
	if err := gocqlx.Select(&inventory, q.Query); err != nil {
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Failed to retrieve inventory")
		return
	}
	// -- End: Retrieve Current Inventory from DB --

	inventoryLookup := make(map[gocql.UUID]*model.InventoryDetails)
	for i := 0; i < len(inventory); i++ {
		inventoryLookup[inventory[i].Id] = &inventory[i]
	}

	// -- Start: Verify Inventory Locally --
	sufficientInventory := true
	type InventoryMutation struct {
		NewQty uint64
		Inventory *model.InventoryDetails
	}
	inventoryMutations := make([]InventoryMutation, 0, len(allocationRequest.Products))
	for _, product := range allocationRequest.Products {
		inv := inventoryLookup[product.Id]
		log.Println(inv)
		if inv == nil {
			log.Println("Unable to find Inventory for Product: ", product.Id)
			sufficientInventory = false
			break
		}
		reqQty, err := strconv.ParseUint(product.Quantity, 10, 64)
		if err != nil {
			log.Println("Received Invalid Product Quantity: ", err)
			sufficientInventory = false
			break
		}
		qty, err := strconv.ParseUint(inv.Quantity, 10, 64)
		if err != nil {
			log.Println("Found Invalid Product Quantity: ", err)
			sufficientInventory = false
			break
		}
		if qty < reqQty {
			log.Println("Insufficient Product Inventory")
			sufficientInventory = false
			break
		}
		inventoryMutations = append(inventoryMutations, InventoryMutation{
			NewQty: qty - reqQty,
			Inventory: inv,
		})
	}

	if !sufficientInventory {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Not Enough Inventory to Allocate")
		return
	}
	// -- End: Verify Inventory Locally

	// -- Start: Mutate Remote Inventory
	batch := qb.Batch()
	batchData := qb.M{}
	for _, inventoryMutation := range inventoryMutations {
		prefix := inventoryMutation.Inventory.Id.String()

		log.Println(inventoryMutation.Inventory)
		batch.AddWithPrefix(prefix,
			qb.Update("inventory").
				Set("quantity").
					Where(qb.Eq("store_id"), qb.Eq("id")))

		batchData[prefix + ".quantity"] = strconv.FormatUint(inventoryMutation.NewQty, 10)
		batchData[prefix + ".store_id"] = inventoryMutation.Inventory.StoreId
		batchData[prefix + ".id"] = inventoryMutation.Inventory.Id
	}
	query, names = batch.ToCql()
	q = gocqlx.Query(dbShard.Query(query), names).BindMap(batchData)
	if err := q.ExecRelease(); err != nil {
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Failed to commit inventory")
		return
	}
	// -- End: Mutate Remote Inventory

	// -- Start: Create Allocation --
	allocationId, err := gocql.RandomUUID()
	if err != nil {
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Unable to allocate new ID")
		return
	}

	updatedProducts := make([]*model.InventoryDetails, 0, len(allocationRequest.Products))
	for _, product := range allocationRequest.Products {
		inv := inventoryLookup[product.Id]
		updatedProducts = append(updatedProducts, &model.InventoryDetails{
			Id: product.Id,
			Name: inv.Name,
			Quantity: product.Quantity,
			Size: inv.Size,
		})
	}

	allocation := model.AllocationDetails{
		UserId: userId,
		Id: allocationId,
		Status: "Unconfirmed",
		Expires: time.Now().Add(time.Minute),
		Products: updatedProducts,
	}

	query, names = qb.Insert("allocations").Columns("user_id", "id", "status", "expires", "products").ToCql()
	q = gocqlx.Query(dbShard.Query(query), names).BindStruct(allocation)
	if err := q.ExecRelease(); err != nil {
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Unable to write allocation")
		return
	}
	// -- End: Create Allocation --

	output.WriteJson(w, allocation)
}

func (ctx *RequestContext) ListAllocations(w http.ResponseWriter, r *http.Request) {
	userId, err := gocql.ParseUUID(common.GetUserId(r))
	if err != nil {
		output.WriteErrorMessage(w, http.StatusUnauthorized, "User Session Not Valid")
		return
	}

	storeId, err := ctx.getStoreIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
		return
	}

	dbShard := SelectShard(storeId, ctx.Database)

	query, names := qb.Select("allocations").Where(qb.Eq("user_id")).ToCql()
	q := gocqlx.Query(dbShard.Query(query), names).BindStruct(model.AllocationDetails{
		UserId: userId,
	})

	allocs := make([]*model.AllocationDetails, 0)
	if err := gocqlx.Select(&allocs, q.Query); err != nil {
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Failed to lookup allocations")
		return
	}

	output.WriteJson(w, model.ListAllocationsResult{
		Allocations: allocs,
	})
}

func (ctx *RequestContext) GetAllocation(w http.ResponseWriter, r *http.Request) {
	userId, err := gocql.ParseUUID(common.GetUserId(r))
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid User Session")
		return
	}

	allocationId, err := getAllocationIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Allocation Identifier")
		return
	}

	storeId, err := ctx.getStoreIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
		return
	}

	dbShard := SelectShard(storeId, ctx.Database)

	// Set up Query
	query, names := qb.Select("allocations").Where(qb.Eq("user_id"), qb.Eq("id")).Limit(1).ToCql()
	q := gocqlx.Query(dbShard.Query(query), names).BindStruct(model.AllocationDetails{
		UserId: userId,
		Id: allocationId,
	})

	// Execute Query
	var allocation model.AllocationDetails
	if err := gocqlx.Get(&allocation, q.Query); err != nil {
		switch err {
		case gocql.ErrNotFound:
			output.WriteErrorMessage(w, http.StatusNotFound, "Allocation not found")
		default:
			log.Println(err)
			output.WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	output.WriteJson(w, allocation)
}

func (ctx *RequestContext) ConfirmAllocation(w http.ResponseWriter, r *http.Request) {
	userId, err := gocql.ParseUUID(common.GetUserId(r))
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid User Session")
		return
	}

	allocationId, err := getAllocationIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Allocation Identifier")
		return
	}

	storeId, err := ctx.getStoreIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
		return
	}

	dbShard := SelectShard(storeId, ctx.Database)

	query, names := qb.Select("allocations").
		Where(qb.Eq("user_id"), qb.Eq("id")).Limit(1).ToCql()
	q := gocqlx.Query(dbShard.Query(query), names).BindStruct(model.AllocationDetails{
		UserId: userId,
		Id: allocationId,
	})

	var allocation model.AllocationDetails
	if err := gocqlx.Get(&allocation, q.Query); err != nil {
		switch err {
		case gocql.ErrNotFound:
			output.WriteErrorMessage(w, http.StatusNotFound, "Store not found")
		default:
			log.Println(err)
			output.WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	if allocation.Status != "Unconfirmed" {
		log.Println("Allocation in unexpected state")
		output.WriteErrorMessage(w, http.StatusBadRequest, "Allocation not in Unconfirmed state")
		return
	}

	allocation.Status = "Confirmed"

	query, names = qb.Update("allocations").
		SetNamed("status", "new_status").
			Where(qb.Eq("user_id"), qb.Eq("id")).
				If(qb.InNamed("status", "old_status")).
					ToCql()
	q = gocqlx.Query(dbShard.Query(query), names).BindMap(qb.M{
		"old_status": []string {"Unconfirmed"},
		"new_status": allocation.Status,
		"user_id": allocation.UserId,
		"id": allocation.Id,
	})

	if err := q.ExecRelease(); err != nil {
		log.Println("Unable to update allocation: ", err)
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Unable to update Allocation")
		return
	}

	output.WriteJson(w, allocation)
}

func (ctx *RequestContext) ExpireAllocation(w http.ResponseWriter, r *http.Request) {
	userId, err := gocql.ParseUUID(common.GetUserId(r))
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid User Session")
		return
	}

	allocationId, err := getAllocationIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Allocation Identifier")
		return
	}

	storeId, err := ctx.getStoreIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
		return
	}

	dbShard := SelectShard(storeId, ctx.Database)

	query, names := qb.Select("allocations").
		Where(qb.Eq("user_id"), qb.Eq("id")).Limit(1).ToCql()
	q := gocqlx.Query(dbShard.Query(query), names).BindStruct(model.AllocationDetails{
		UserId: userId,
		Id: allocationId,
	})

	var allocation model.AllocationDetails
	if err := gocqlx.Get(&allocation, q.Query); err != nil {
		switch err {
		case gocql.ErrNotFound:
			output.WriteErrorMessage(w, http.StatusNotFound, "Store not found")
		default:
			log.Println(err)
			output.WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	if allocation.Status != "Unconfirmed" {
		log.Println("Allocation in unexpected state")
		output.WriteErrorMessage(w, http.StatusBadRequest, "Allocation not in Unconfirmed state")
		return
	}

	allocation.Status = "Expired"

	query, names = qb.Update("allocations").
		SetNamed("status", "new_status").
		Where(qb.Eq("user_id"), qb.Eq("id")).
		If(qb.InNamed("status", "old_status")).
		ToCql()
	q = gocqlx.Query(dbShard.Query(query), names).BindMap(qb.M{
		"old_status": []string {"Unconfirmed"},
		"new_status": allocation.Status,
		"user_id": allocation.UserId,
		"id": allocation.Id,
	})

	if err := q.ExecRelease(); err != nil {
		output.WriteErrorMessage(w, http.StatusInternalServerError, "")
		return
	}

	output.WriteJson(w, allocation)
}

// -- Helper Functions --

func getAllocationIdFromRequest(r *http.Request) (gocql.UUID, error) {
	return gocql.ParseUUID(mux.Vars(r)["allocation_id"])
}