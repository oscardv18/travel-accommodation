package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "admin"
	password = "secret"
	dbname   = "travel"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var db, err = sql.Open("postgres", psqlInfo)

func Init() {
	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Database Connected successfully! :)")
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
