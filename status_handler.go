package main

import (
	"encoding/json"
	"net/http"
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// StatusUpdate updates light status based on radio button click
func StatusUpdate(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(getStubbedStatus())
	if err != nil {
		handleInternalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	writeJsonResponse(w, json)
}

func getStubbedStatus() Status {
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/app")
if err != nil {
panic(err.Error())
}
defer db.Close()

  var (
    id int
    queue string
    status string
    )
    rows, err := db.Query("select * from light_status where id = ?", 5)
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
		Queue: 	queue,
		Status:    status,
	}
}