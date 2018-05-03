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
	"errors"
)

// -- Request Handlers --
func (ctx *RequestContext) ListStores(w http.ResponseWriter, r *http.Request) {
	// Set up Query
	query, names := qb.Select("stores").ToCql()
	storeChan := make(chan[]model.StoreDetails)
	getFunc := func(session *gocql.Session, storeChan chan[]model.StoreDetails) {
		q := gocqlx.Query(ctx.Database.Shard1.Query(query), names)

		// Execute Query
		stores := make([]model.StoreDetails, 0)
		if err := gocqlx.Iter(q.Query).Select(&stores); err != nil {
			log.Println(err)
			storeChan <- nil
		} else {
			storeChan <- stores
		}
	}
	go getFunc(ctx.Database.Shard1, storeChan)
	go getFunc(ctx.Database.Shard2, storeChan)

	stores := make([]model.StoreDetails, 0)
	for i := 0; i < 2; i++ {
		stores = append(stores, <-storeChan...)
	}

	// Transform Output to JSON
	output.WriteJson(w, &model.ListStoresResult{
		Stores:        stores,
		NextPageToken: nil,
	})
}

func (ctx *RequestContext) GetStore(w http.ResponseWriter, r *http.Request) {
	storeId, err := ctx.getStoreIdFromRequest(r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
		return
	}

	dbShard := SelectShard(storeId, ctx.Database)

	// Set up Query
	query, names := qb.Select("stores").Where(qb.Eq("id")).Limit(1).ToCql()
	q := gocqlx.Query(dbShard.Query(query), names).BindStruct(model.StoreDetails{
		Id: storeId,
	})

	// Execute Query
	var store model.StoreDetails
	if err := gocqlx.Get(&store, q.Query); err != nil {
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
func (ctx *RequestContext) getStoreIdFromRequest(r *http.Request) (gocql.UUID, error) {
	return ctx.getStoreId(mux.Vars(r)["store_id"])
}

func (ctx *RequestContext) getStoreId(rawStoreId string) (gocql.UUID, error) {
	storeId, err := gocql.ParseUUID(rawStoreId)
	if err != nil {
		storeId, err = ctx.getStoreIdByName(rawStoreId)
		if err != nil {
			return gocql.UUID{}, err
		}
	}
	return storeId, nil
}

func (ctx *RequestContext) getStoreIdByName(name string) (gocql.UUID, error) {
	query, names := qb.Select("stores").Columns("id").Where(qb.Eq("name")).Limit(1).ToCql()
	storeChan := make(chan *model.StoreDetails)
	getFunc := func(session *gocql.Session, c chan*model.StoreDetails) {
		q := gocqlx.Query(session.Query(query), names).BindStruct(model.StoreDetails{
			Name: name,
		})

		var storeDetails model.StoreDetails
		if err := gocqlx.Get(&storeDetails, q.Query); err != nil {
			log.Println(err)
			storeChan <- nil
		} else {
			storeChan <- &storeDetails
		}
	}
	go getFunc(ctx.Database.Shard1, storeChan)
	go getFunc(ctx.Database.Shard2, storeChan)

	var result *model.StoreDetails
	var count uint32
	for result == nil && count < 2 {
		result = <-storeChan
		count++
	}

	if result == nil {
		return gocql.UUID{}, errors.New("couldn't find store")
	} else {
		return result.Id, nil
	}
}