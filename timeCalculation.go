package main 

import (
	"fmt"
	"time"
) 

func TimeCalculation(){
	start := time.Now()
    /* Your Buisness logic*/

	sum := 0 

	for i := 1 ; i < 100000000 ; i++{
		sum+= i*i
	}
	fmt.Println(sum)
	end := time.Since(start)
	fmt.Println(end.Milliseconds())
}
