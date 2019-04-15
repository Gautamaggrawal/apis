package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Handle API routes
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/adduser", userAddHandler).Methods("POST")
	api.HandleFunc("/getusers", returnAllUsers).Methods("GET")
	api.HandleFunc("/search", emailsearchHandler).Methods("POST")

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Println("http://localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
