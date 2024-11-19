package main

import (
	"net/http"
)

func htmlServer(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"index.html")
}

func index(){
	http.HandleFunc("/",htmlServer)
	_ = http.ListenAndServe(":8080",nil)

}