package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func connect(){
	dbString := "root:Karthik@tcp(127.0.0.1:3306)/c5i"
	db,err := sql.Open("mysql",dbString)

	if err!=nil{
		log.Fatal("error in string")
	}
	defer db.Close()

	err = db.Ping()

	if err!= nil{
		log.Fatal("Unable to connect to db")
	}

	fmt.Println("Sucessfully connected to databse")
}