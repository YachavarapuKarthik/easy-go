package main

import "fmt"

func calci() {
    var a int
	var b  int 
	fmt.Scan(&a)
	fmt.Scan((&b))
    fmt.Println("Addition is ",a+b)
	fmt.Println("Subtraction is ",a-b)
	fmt.Println("Multiplication is",a*b)
	fmt.Println("Subraction is",a/b)
	fmt.Println("Modulus is: ",a%b)
}
