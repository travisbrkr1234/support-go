package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  )

<<<<<<< HEAD
  func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/app")
=======
  func pingDatabase() {
    db, err := sql.Open("mysql", "root:SOMEPASSWORD@tcp(127.0.0.1:3306)/app")
>>>>>>> 4dbdab67b87e407d353b1c0929d3cd5f467e4b2a
  if err != nil {
  panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
  }
  defer db.Close()

  // Open doesn't open a connection. Validate DSN data:
  err = db.Ping()
  if err != nil {
  panic(err.Error()) // proper error handling instead of panic in your app
  }

  fmt.Printf("Great success!");
}
