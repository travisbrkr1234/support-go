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

			var (
				id int
				status string
				)
				rows, err := db.Query("select id, status from light_status where id = ?",7)
				if err != nil {
					panic(err)
				}
				defer rows.Close()
				for rows.Next() {
					err := rows.Scan(&id, &status)
					if err != nil {
						panic(err)
					}
					fmt.Println(id, status)
				}
				err = rows.Err()
				if err != nil {
					panic(err)
				}
}
