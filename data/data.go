package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "default"
	password = "pass"
	dbname   = "Data"
)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)
	// defer db.Close()

	err = db.Ping()
	checkError(err)

	fmt.Println("Connected to database")
	return db
}

func SetupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	checkError(err)

	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
