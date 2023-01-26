package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// database details
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "H@rsh@130"
	dbname   = "BloggerMvc"
)


// global variable
var DB *sql.DB

// database connection is created here
func ConnectDataBase() {
	var err error

	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")
}
