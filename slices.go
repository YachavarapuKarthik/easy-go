package main 

import "fmt"

func slice(){
	//var arr_slice []int
	arr_slice :=[]int{}
	for i:=0 ; i < 20 ; i++{
		arr_slice = append(arr_slice,i+1)
		fmt.Print(arr_slice[i]," ")
	}
	fmt.Println("\n",len(arr_slice))
	fmt.Print(arr_slice[:2])
}