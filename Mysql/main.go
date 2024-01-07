package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "testdb",
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Println("Db connection is success")
	defer db.Close()

}
