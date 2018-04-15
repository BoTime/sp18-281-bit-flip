package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

type OutputHelper struct {
	w http.ResponseWriter
}

func (out OutputHelper) WriteErrorMessage(status int, message string) {
	out.w.WriteHeader(status)
	fmt.Fprintf(out.w, message)
}

func (out OutputHelper) WriteJson(obj interface{}) {
	if output, err := json.Marshal(obj); err != nil {
		log.Println(err)
		out.WriteErrorMessage(http.StatusInternalServerError, "Internal Server Error")
		return
	} else {
		out.w.Header().Set("content-type", "application/json")
		out.w.WriteHeader(http.StatusOK)
		out.w.Write(output)
	}
}