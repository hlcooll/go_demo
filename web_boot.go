package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/"，requestHandler)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func requestHandler(w http.ResponseWriter ,r *http.Request){
     if r.Method == "POST" {
		 body, err = ioutil.ReadAll(r.Body)

	 }
}