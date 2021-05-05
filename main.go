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
	password = "ZA21122002"
	dbname   = "zybair"
)

func main() {
	a := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbname)
	db, err := sql.Open("postgres", a)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
