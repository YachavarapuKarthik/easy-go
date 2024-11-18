package main 

import (
	"fmt"
	"net/http"
)

func helperFunction(w http.ResponseWriter, r *http.Request){
	msg := "Hey karthik WEb server is started"
	fmt.Fprintln(w,msg)
}

func StartServer(){
	http.HandleFunc("/",helperFunction)
	err:= http.ListenAndServe(":8080",nil)
	if err!= nil{
		fmt.Println("Unable to start web server")
	}
}