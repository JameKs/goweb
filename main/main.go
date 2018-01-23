package main

import (
	"fmt"
	"net/http"
	"log"
)
func main() {
	fmt.Print("Hello World!")
	http.HandleFunc("/",sayhelloName)
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal("ListenAndService",err)
	}
}

func sayhelloName(w http.ResponseWriter,r *http.Request)  {
r.ParseForm()
fmt.Print(r.Form)
fmt.Printf("path",r.URL.Path)
fmt.Printf("scheme",r.URL.Scheme)
fmt.Print(r.Form["url_long"])
	for k,v := range r.Form  {
		fmt.Println("key",k)
		fmt.Printf("val",v)
	}
	fmt.Fprintln(w,"Hello World!")
}