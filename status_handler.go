package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

// CurrentStatus gets current light status for all queues
// StatusUpdate updates light status based on radio button click
func CurrentStatus(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(getStubbedStatuses())
	if err != nil {
		handleInternalServerError(w, err)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	writeJsonResponse(w, json)
}

func getStubbedStatuses() Statuses {
	var (
		id    int
		queue string
		color string
	)
	rows, err := DB.Query("select * from light_status")
	if err != nil {
		panic("Failed to Login")
	}
	var status Status
	statuses := Statuses{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &queue, &color)
		if err != nil {
			panic(err)
		}
		status = Status{
			ID:     id,
			Queue:  queue,
			Status: color,
		}
		statuses.Data = append(statuses.Data, status)
		fmt.Println(id, queue, status)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return statuses
}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	queue := mux.Vars(r)["queue"]
	//w.Write([]byte(fmt.Sprintf(queue)))
	fmt.Println(queue)
	var status Status
	statusDecoder := json.NewDecoder(r.Body)
	if err := statusDecoder.Decode(&status); err != nil {
		fmt.Println("Unable to marshal json to Status")
		return
	}
	status.Queue = queue
	fmt.Printf("Status: %s\n", status.Status)
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtUpd, err := db.Prepare("UPDATE light_status SET status = ? WHERE queue = ?")
	if err != nil {
		panic(err)
	}

	res, err := stmtUpd.Exec(status.Status, status.Queue)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", affect)
}

type Statuses struct {
	Data []Status `json:"data"`
}
