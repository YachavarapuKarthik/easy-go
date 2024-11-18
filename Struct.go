package main 

import "fmt"

type Person struct{
	name string
	age int 
	id int 
}

func PersonImp(){
	karthik := Person{name:"karthik",age:20,id:65}
	fmt.Println(karthik)
}