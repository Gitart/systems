package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connect to mysql now")
	db, _ := sql.Open("mysql", "account:password@tcp(host:3306)/db")
	defer db.Close()
	var sku_name string
	results, _ := db.Query("sql_cmd")
	defer db.Close()
	for results.Next() {
		results.Scan(&sku_name)
		fmt.Println(sku_name)
	}
}