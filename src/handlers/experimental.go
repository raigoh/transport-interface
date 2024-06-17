package handlers

import (
	"net/http"
)

func ExperimentalHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Experimental Endpoint"))
}
