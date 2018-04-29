package handler

import (
	"cmpe281/common/output"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
	"cmpe281/inventory/model"
	"log"
	"net/http"
)

func (ctx *RequestContext) GetInventory(w http.ResponseWriter, r *http.Request) {
	storeId, err := ctx.getStoreId(w, r)
	if err != nil {
		output.WriteErrorMessage(w, http.StatusBadRequest, "Invalid Store Identifier")
	}

	// Set up Query
	query, names := qb.Select("products").Where(qb.Eq("store_id")).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindStruct(model.ProductDetails{
		StoreId: storeId,
	})

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
