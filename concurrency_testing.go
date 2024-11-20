package main  

import (
	"fmt"
	"time"
	"sync"
)

var waitGroups = sync.WaitGroup{}

func miniTask1(){
	time.Sleep( 10* time.Second)
	fmt.Println("MiniTask1 is completed")
	waitGroups.Done()
}
func miniTask2(){
	time.Sleep( 20* time.Second)
	fmt.Println("MiniTask2 is completed")
	waitGroups.Done()
}
func mainTask(){
	fmt.Println("Hey main task is running ")
	waitGroups.Add(2)
	go miniTask1()
	go miniTask2()
	waitGroups.Wait()

}