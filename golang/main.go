package main

import (
	"fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
) 

func main () { 

	router:= mux.NewRouter()
	router.HandleFunc("/add", userAddHandler).Methods("POST") 	
	router.HandleFunc("/getproducts", returnAllUsers).Methods("GET") 
	http.Handle("/", router) 
	fmt.Println("Connected to port 1234") 
	// log.Fatal(http.ListenAndServe (" : 1234 ", router)) 
	log.Fatal(http.ListenAndServe(":12346", router))

}

