package output

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

func WriteErrorMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, message)
}

func WriteJson(w http.ResponseWriter, obj interface{}) {
	if output, err := json.Marshal(obj); err != nil {
		log.Println(err)
		WriteErrorMessage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	} else {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	}
}