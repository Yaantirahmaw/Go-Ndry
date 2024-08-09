package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = ""
	password = ""
	dbname = "go-ndry"
)

var PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}else{
		fmt.Println("CONNECTED!")
	}
	return db
}