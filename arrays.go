package main 

import "fmt"

func arrays(){
	var array [50]int
	for i:=0 ; i<len(array); i++{
		array[i] = i+1 
	}

	for i:=0 ; i<len(array) ; i++{
		fmt.Print(array[i], " ")
	}
 
	for _,val := range array {
		fmt.Printf("Index: %d\n", val)
	}
	
}