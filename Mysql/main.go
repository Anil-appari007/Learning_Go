package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := os.Getenv("DBUSER")
	dbPasswd := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")
	if dbUser == "" || dbPasswd == "" || dbName == "" {
		log.Fatal("Check the Database Creds")
	} else {
		log.Println("Database Credentials Found & not empty")
	}
	config := mysql.Config{
		User:   dbUser,
		Passwd: dbPasswd,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: dbName,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Db Connection is success")

	getData, err := db.Query("SELECT * FROM cities;")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Retreiving Data from table")
	for getData.Next() {
		var id int
		var city string
		var population int
		err = getData.Scan(&id, &city, &population)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("%v. City: %s\t-Population: %v", id, city, population)
	}

	defer getData.Close()

}
