package main

import "fmt"

var name string
var id int
var age int

func human(){
	fmt.Print("Enter name, id , age")
	fmt.Scan(&name)
	fmt.Scan(&id)
	fmt.Scan(&age)
	fmt.Print(name," ",id," ",age)
}