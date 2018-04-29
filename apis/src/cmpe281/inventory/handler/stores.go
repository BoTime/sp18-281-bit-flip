package handler

import (
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
	"cmpe281/inventory/model"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"cmpe281/common/output"
	"log"
)

// -- Request Handlers --
func (ctx *RequestContext) ListStores(w http.ResponseWriter, r *http.Request) {
	// Set up Query
	query, names := qb.Select("stores").ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names)

	// Execute Query
	stores := make([]model.StoreDetails, 0)
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
	store, err := ctx.getStoreById(w, r)
	if err != nil {
		store, err = ctx.getStoreByName(w, r)
	}

	if err != nil {
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

// -- Helper Functions --
func (ctx *RequestContext) getStoreId(w http.ResponseWriter, r *http.Request) (gocql.UUID, error) {
	vars := mux.Vars(r)

	storeId, err := gocql.ParseUUID(vars["store_id"])
	if err != nil {
		storeId, err = ctx.getStoreIdByName(vars["store_id"])
		if err != nil {
			return gocql.UUID{}, err
		}
	}

	return storeId, nil
}

func (ctx *RequestContext) getStoreIdByName(name string) (gocql.UUID, error) {
	query, names := qb.Select("stores").Columns("id").Where(qb.Eq("name")).Limit(1).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindStruct(model.StoreDetails{
		Name: name,
	})

	var store model.StoreDetails
	if err := gocqlx.Get(&store, q.Query); err != nil {
		return gocql.UUID{}, err
	}

	return store.Id, nil
}

func (ctx *RequestContext) getStoreById(w http.ResponseWriter, r *http.Request) (*model.StoreDetails, error) {
	// Parse UUID
	storeId, err := gocql.ParseUUID(mux.Vars(r)["store_id"])
	if err != nil {
		return nil, err
	}

	// Set up Query
	query, names := qb.Select("stores").Where(qb.Eq("id")).Limit(1).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindStruct(model.StoreDetails{
		Id: storeId,
	})

	// Execute Query
	var store model.StoreDetails
	if err := gocqlx.Get(&store, q.Query); err != nil {
		return nil, err
	}

	return &store, nil
}

func (ctx *RequestContext) getStoreByName(w http.ResponseWriter, r *http.Request) (*model.StoreDetails, error) {
	// Secondary Index allows us to lookup ID by Store Name

	// Set up Query
	query, names := qb.Select("stores").Where(qb.Eq("name")).Limit(1).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindStruct(model.StoreDetails{
		Name: mux.Vars(r)["store_id"],
	})

	// Execute Query
	var store model.StoreDetails
	if err := gocqlx.Get(&store, q.Query); err != nil {
		return nil, err
	}

	return &store, nil
}