package handler

import (
	"cmpe281/common/output"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
	"cmpe281/inventory/model"
	"log"
	"net/http"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func (ctx *RequestContext) GetInventory(w http.ResponseWriter, r *http.Request) {
	storeId, err := ctx.getStoreIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
	}

	// Set up Query
	query, names := qb.Select("inventory").Where(qb.Eq("store_id")).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindStruct(model.InventoryDetails{
		StoreId: storeId,
	})

	// Execute Query
	products := make([]model.InventoryDetails, 0)
	if err := gocqlx.Iter(q.Query).Select(&products); err != nil {
		log.Println(err)
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Transform Output to JSON
	output.WriteJson(w, &model.ListInventoryResult{
		Products: products,
	})
}

// -- Helper Functions --
func (ctx *RequestContext) getProductIdFromRequest(r *http.Request) (gocql.UUID, error) {
	return ctx.getProductId(mux.Vars(r)["product_id"])
}

func (ctx *RequestContext) getProductId(rawProductId string) (gocql.UUID, error) {
	productId, err := gocql.ParseUUID(rawProductId)
	if err != nil {
		productId, err = ctx.getProductIdByName(rawProductId)
		if err != nil {
			return gocql.UUID{}, err
		}
	}

	return productId, nil
}

func (ctx *RequestContext) getProductIdByName(name string) (gocql.UUID, error) {
	query, names := qb.Select("products").Distinct("id").Where(qb.Eq("name")).Limit(1).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindStruct(model.ProductDetails{
		Name: name,
	})

	var product model.ProductDetails
	if err := gocqlx.Get(&product, q.Query); err != nil {
		return gocql.UUID{}, err
	}

	return product.Id, nil
}