package main 

import "fmt"

func slice(){
	var arr_slice []int
	for i:=0 ; i < 20 ; i++{
		arr_slice = append(arr_slice,i+1)
		fmt.Print(arr_slice[i]," ")
	}

}