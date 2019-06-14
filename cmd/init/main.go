package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("mysql schema creation")

	db, err := sql.Open("mysql", "entities:entities@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic("cannot connect: " + err.Error())
	}

	defer db.Close()

	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS entities")
	if err != nil {
		panic("cannot create db: " + err.Error())
	}

	db, err = sql.Open("mysql", "entities:entities@tcp(127.0.0.1:3306)/entities")
	if err != nil {
		panic("cannot connect: " + err.Error())
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS companies (
		id INT UNSIGNED AUTO_INCREMENT NOT NULL,
		name VARCHAR(100) NOT NULL,
		town VARCHAR(100) CHARACTER SET 'ascii' NOT NULL,
		postcode_in VARCHAR(4) CHARACTER SET 'ascii' NOT NULL,
		postcode_out CHAR(3) CHARACTER SET 'ascii' NOT NULL,
		uksic_code CHAR(5) CHARACTER SET 'ascii' NULL,
		PRIMARY KEY (id));`)
	if err != nil {
		panic("cannot create db: " + err.Error())
	}

	// insert, err := db.Query(`INSERT into companies (name, town, postcode_in, postcode_out, uksic_code) VALUES( ?, ?, ?, ?, ? )`, "BOUJELLÉ LTD", "LONDON", "HA8", "6LY", "46450")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
}
