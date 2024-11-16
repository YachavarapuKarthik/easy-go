package main 

import "fmt"

func conditional(){
	var x int
	fmt.Scan(&x)
	if x==0{
		fmt.Println("X is 0")
	}else if x % 2 == 0 {
		fmt.Println("X is even")
	}else{
		fmt.Println("X is odd")
	}
}