package handler

import (
	"cmpe281/common/output"
	"net/http"
)

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
