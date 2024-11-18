package main

import "fmt"

func MapCreation(){
	fmt.Println("Creating a map")
	// normal map 
	// var maps map[string]int
	//
	// List of maps can be createdd as 
	// var maps []map[string]int
	// 
	// Dynamic map creation is
	// mini_map := map[string]int


	var maps []map[string]int 

	mini_map := map[string]int{"Maths":10,"English":20}
	
	maps = append(maps,mini_map)

	fmt.Println("List of maps are :")
	for key,value := range maps{
		fmt.Println("Key: ",key," Map: ",value)
	}

	fmt.Println("Single map :")

	for key,value := range mini_map{
		fmt.Println("Key:",key," Map: ",value)
	}
}