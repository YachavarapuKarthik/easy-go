package main 

import "fmt"

func calculator(){
	add := func(a int , b int) int{
		return a+b
	}

	sub := func(a int, b int ) int {
		return a-b
	}

	fmt.Println(add(10,20))
	fmt.Println(sub(10,20))

}