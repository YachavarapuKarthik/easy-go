package main

import (
	"fmt"
	"net/http"
)

func htmlServer(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"index.html")
}

func index(){
	http.HandleFunc("/",htmlServer)
	err := http.ListenAndServe(":8080",nil)

	if err != nil{
		fmt.Println("Errror")
	}

}