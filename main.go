package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// employee represents data about a employee.
type employee struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address address `json:"address"`
	Salary  float64 `json:"salary"`
}

type address struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Pincode int    `json:"pincode"`
}

//address

var location = []address{
	{City: "Vijayawada", Pincode: 521190, Country: "India"},
	{City: "Pune", Pincode: 986134, Country: "India"},
	{City: "Adelaide", Pincode: 7910, Country: "Australia"},
	{City: "Toronto", Pincode: 6213, Country: "Canada"},
}

// employees slice to seed record album data.
var employees = []employee{
	{ID: "1", Name: "Ram", Address: location[0], Salary: 50000.00},
	{ID: "2", Name: "Harish", Address: location[1], Salary: 50220.00},
	{ID: "3", Name: "Lakshmi", Address: location[2], Salary: 20000.00},
	{ID: "4", Name: "Vamshi", Address: location[3], Salary: 80000.56},
}

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", getEmployeeByID)
	router.POST("/employees", postEmployees)
	router.PUT("/employees/:id", updateEmployeeByID)
	router.DELETE("/employees/:id", deleteEmployeeById)
	router.Run("localhost:8080")
}

func RemoveIndex(s []employee, index int) []employee {
	return append(s[:index], s[index+1:]...)
}

func deleteEmployeeById(c *gin.Context) {
	id := c.Param("id")
	for i, data := range employees {
		if data.ID == id {
			employees = RemoveIndex(employees, i)
			return
		}
	}
}

// update the employee details with list of employees information
func updateEmployeeByID(c *gin.Context) {
	var newEmployee employee
	id := c.Param("id")

	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}

	// Loop through the list of employees, looking for
	// an employee whose ID value matches the parameter.
	for i, a := range employees {
		if a.ID == id {
			employees[i] = newEmployee
			c.IndentedJSON(http.StatusAccepted, newEmployee)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// getEmployees details with list of employees information
func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

// postEmployees create a new employee with json body posted in the call
func postEmployees(c *gin.Context) {
	var newEmployee employee

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}

	// Add the new employee to the slice.
	employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

// getEmployeeByID locates the employee whose ID value matches the id
// parameter sent by the client, then returns that employee as a response.
func getEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of employees, looking for
	// an employee whose ID value matches the parameter.
	for _, a := range employees {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
}
