package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	create("goTest")
}

func create(name string) {
	db, err := sql.Open("mysql", "root:password@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE test(id integer, data varchar(32))")
	if err != nil {
		panic(err)
	}
}
