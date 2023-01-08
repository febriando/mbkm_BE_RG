package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	// // menutup db
	// defer db.Close()

	// // melakukan ping ke database PostgreSQL
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// create table employee
	// _, err = db.Exec(`CREATE TABLE employee (
	// 	id INT,
	// 	name VARCHAR(255) NOT NULL,
	// 	age INT NOT NULL,
	// 	address VARCHAR(255),
	// 	salary INT
	//   )`)
	// fmt.Println("Table employee created")
	// rename table employee to employees

	_, err = db.Exec(`ALTER TABLE employee RENAME TO employees`)
	fmt.Println("Table employee RENAME to employees")

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
