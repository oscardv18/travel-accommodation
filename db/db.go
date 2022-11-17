package db

import (
	"database/sql"
	"fmt"
)

// postgresql db connection
var psqlconn = "postgresql://admin:secret@postgres/travel?sslmode=disable"

func Init() {
	// open db
	var db, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	// close db
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Database Connected successfully! :)")
}

func CheckError(err error) {
	if err != nil {
		e := fmt.Errorf("error at save record %d", 3312)
		// log.Fatal(err)
		fmt.Println(e)
	}
}
