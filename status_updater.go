package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func StatusUpdate(w http.ResponseWriter, r *http.Request) {
	queue := mux.Vars(r)["/{queue}"]
	w.Write([]byte(fmt.Sprintf(queue)))
}
