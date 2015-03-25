package main

import (
  "bufio"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  )

		func main() {
      consolereader := bufio.NewReader(os.Stdin)
      fmt.Println("Enter a number")
      fmt.Println("1 for Phone")
      fmt.Println("3 for Chat")
      fmt.Println("5 for Tier2")
      fmt.Println("7 for Online")
      input, err := consolereader.ReadString('\n') // this will prompt the user for input

      if err != nil {
              fmt.Println(err)
              os.Exit(1)
      }

			db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/app")
	  if err != nil {
	  panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	  }
	  defer db.Close()

			var (
				id int
				status string
				)
				rows, err := db.Query("select id, status from light_status where id = ?", input)
				if err != nil {
					panic("Failed to Login")
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
