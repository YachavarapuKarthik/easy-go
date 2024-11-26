package main

import (
	"fmt"
)

func errorThroowing() error{
	return fmt.Errorf("Hii welcome errooros🙂")
}

func errThrow(){
	err := errorThroowing()
	if err != nil{
		fmt.Println("Error is ",err)
	}
}