package main 

import(
	"fmt"
	"easygo/helper"
	"time"
)

func isValidAge(){
	start := time.Now()
	var age int 
	fmt.Println("Enter Applicant age")
	fmt.Scan(&age)
	if helper.AgeValidator(age){
		fmt.Println("Can able to vote")
	}else {
		fmt.Println("Cannot able to vote")
	}
	end := time.Since(start)
	fmt.Println(end)
}
