package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/app")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtUpd, err := db.Prepare("UPDATE light_status SET status = ?, WHERE queue = ?")
	if err != nil {
		panic(err)
	}

	res, err := stmtUpd.Exec("test", 9)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(affect)
}
