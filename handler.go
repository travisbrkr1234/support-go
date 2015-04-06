package main

import (
	"net/http"
)

func handleInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
func writeJsonResponse(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
