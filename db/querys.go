package db

import (
	"fmt"
)

// insert data
func InsertUserData(fname string, lname string, email string, pnumber uint64) {
	insertDynStmt := `INSERT INTO users(firstname,lastname,email,phonenumber) VALUES($1,$2,$3,$4)`
	_, e := db.Exec(insertDynStmt, fname, lname, email, pnumber)
	CheckError(e)
	fmt.Println("New record saved in db")
}
