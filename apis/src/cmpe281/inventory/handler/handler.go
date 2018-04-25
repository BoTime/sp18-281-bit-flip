package handler

import (
	"cmpe281/common/output"
	"github.com/gocql/gocql"
	"net/http"
)

type RequestContext struct {
	Cassandra *gocql.Session
}

func (ctx *RequestContext) ListStores(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
}

func (ctx *RequestContext) GetStore(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
}

func (ctx *RequestContext) GetInventory(w http.ResponseWriter, r *http.Request) {
	output.WriteErrorMessage(w, http.StatusNotImplemented, "To be implemented...")
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
