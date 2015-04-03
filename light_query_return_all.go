package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  )

		func main() {
      db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/app")
	  if err != nil {
	  panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	  }
	  defer db.Close()

			var (
				id int
        queue string
				status string
				)
				rows, err := db.Query("select * from light_status")
				if err != nil {
					panic("Failed to Login")
				}
				defer rows.Close()
				for rows.Next() {
					err := rows.Scan(&queue, &status, &id)
					if err != nil {
						panic(err)
					}
					fmt.Println(queue, status, id)
				}
				err = rows.Err()
				if err != nil {
					panic(err)
				}
}
