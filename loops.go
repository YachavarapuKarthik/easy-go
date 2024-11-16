package main

import "fmt"

func loops() {
	// for {
	// 	fmt.Println("This is infinite loop")
	// }

	// This acts like a while loop
	// i:=0

	// for i<=5{
	// 	fmt.Println(i)
	// 	i++
	// }

	// Standard for loop

	// for i:=0 ; i<=5 ; i++{
	// 	fmt.Println(i)
	// }


	// continue and break 

	for i:=0 ; i<10 ; i++{
		if i==2{
			continue
		}else if i==9 {
			break
		}else{
			fmt.Println(i)
		}
	}

}
