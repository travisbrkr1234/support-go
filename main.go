package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dbURL string

func main() {
	port := os.Getenv("PORT")
	dbURL = os.Getenv("DATABASE_URL")
	if len(dbURL) == 0 {
		dbURL = "root:@tcp(127.0.0.1:3306)/app"
	}

	DB = getDB(dbURL)
	defer DB.Close()

	fmt.Println("listening...")
	r := NewLoggedRouter()
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}

}

func getDB(dbURL string) *sql.DB {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("Unable to establish a DB connection with url ", dbURL)
		log.Fatal(err)
	}
	return db
}
