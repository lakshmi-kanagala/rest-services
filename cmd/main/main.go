package main

import (
	"fmt"
	"log"
	"net/http"

	"rest-services/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/employees/{id}", services.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees/", services.DeleteEmployees).Methods("DELETE")
	router.HandleFunc("/employees", services.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", services.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", services.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", services.UpdateEmployees).Methods("PUT")
	fmt.Println("Server at 8080")
	http.ListenAndServe(":9000", router)
	log.Fatal(http.ListenAndServe(":9000", router))

}

func main() {
	initializeRouter()

}
