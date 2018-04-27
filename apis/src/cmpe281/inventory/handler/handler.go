package handler

import (
	"cmpe281/common/output"
	"cmpe281/inventory/model"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"log"
	"net/http"
)

type RequestContext struct {
	Cassandra *gocql.Session
}

func (ctx *RequestContext) ListStores(w http.ResponseWriter, r *http.Request) {
	// Set up Query
	query, names := qb.Select("stores").ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names)

	// Execute Query
	var stores []model.StoreDetails
	if err := gocqlx.Iter(q.Query).Select(&stores); err != nil {
		log.Println(err)
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Build Next Page Token
	var nextPageToken *gocql.UUID
	if len(stores) > 0 {
		nextPageToken = &stores[len(stores)-1].Id
	}

	// Transform Output to JSON
	output.WriteJson(w, &model.ListStoresResult{
		Stores:        stores,
		NextPageToken: nextPageToken,
	})
}

func (ctx *RequestContext) GetStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	querySelectors := qb.M{
		"id": nil,
	}

	// Parse Query Inputs
	if storeId, err := gocql.ParseUUID(vars["store_id"]); err == nil {
		querySelectors["id"] = storeId
	} else {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Store ID not provided")
		return
	}

	// Set up Query
	query, names := qb.Select("stores").Where(qb.Eq("id")).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindMap(querySelectors)

	// Execute Query
	var store model.StoreDetails
	if err := gocqlx.Iter(q.Query).Get(&store); err != nil {
		switch err {
		case gocql.ErrNotFound:
			output.WriteErrorMessage(w, http.StatusNotFound, "Store not found")
		default:
			log.Println(err)
			output.WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	output.WriteJson(w, store)
}

func (ctx *RequestContext) GetInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	querySelectors := qb.M{
		"store_id": nil,
	}

	// Parse Query Inputs
	if storeId, err := gocql.ParseUUID(vars["store_id"]); err == nil {
		querySelectors["store_id"] = storeId
	} else {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Store ID not provided")
		return
	}

	// Set up Query
	query, names := qb.Select("products").Where(qb.Eq("store_id")).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindMap(querySelectors)

	// Execute Query
	var products []model.ProductDetails
	if err := gocqlx.Iter(q.Query).Select(&products); err != nil {
		log.Println(err)
		output.WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Transform Output to JSON
	output.WriteJson(w, &model.ListProductsResult{
		Products: products,
	})
}

func (ctx *RequestContext) CreateAllocation(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
}

func (ctx *RequestContext) ListAllocations(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
}

func (ctx *RequestContext) GetAllocation(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
}

func (ctx *RequestContext) ConfirmAllocation(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
}

func (ctx *RequestContext) DeleteAllocation(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
}
