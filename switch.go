package main

import "fmt"

func switch_statements(){
	fmt.Println(" Enter your number")
	var num int
	fmt.Scan(&num)
	switch num{
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("This is default code , enter from 1 to 3 only")
	}
}