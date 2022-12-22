package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/employees/{id}", DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees/", DeleteEmployees).Methods("DELETE")
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/employees", createEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", UpdateEmployees).Methods("PUT")
	fmt.Println("Server at 8080")
	http.ListenAndServe(":9000", router)
	log.Fatal(http.ListenAndServe(":9000", router))

}

func main() {
	initializeRouter()

}
