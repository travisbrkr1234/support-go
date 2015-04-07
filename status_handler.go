package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CurrentStatus gets current light status for all queues
// StatusUpdate updates light status based on radio button click
func CurrentStatus(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(getStubbedStatus())
	if err != nil {
		handleInternalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	writeJsonResponse(w, json)
}

func getStubbedStatus() Status {
	var (
		id     int
		queue  string
		status string
	)
	rows, err := DB.Query("select * from light_status where id = ?", 7)
	if err != nil {
		panic("Failed to Login")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &queue, &status)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, queue, status)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return Status{
		ID:     id,
		Queue:  queue,
		Status: status,
	}
}
