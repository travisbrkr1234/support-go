package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  )

  func main() {
    db, err := sql.Open("mysql", "root:SOMEPASSWORD@tcp(127.0.0.1:3306)/app")
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
