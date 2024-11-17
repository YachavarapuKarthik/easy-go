package main 

import(
	"fmt"
	"easygo/helper"
)

func isValidAge(){
	var age int 
	fmt.Println("Enter Applicant age")
	fmt.Scan(&age)
	if helper.AgeValidator(age){
		fmt.Println("Can able to vote")
	}else {
		fmt.Println("Cannot able to vote")
	}
}
