package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"rest-services/utils"

	"github.com/gorilla/mux"

	"rest-services/model"
	"rest-services/repository"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	// var response []JsonResponse
	var employees []model.Employee
	rows := repository.GetEmployees()
	// Foreach movie
	for rows.Next() {
		var employeeId int
		var employeeName string
		var projectName string
		var salary int

		err := rows.Scan(&employeeId, &employeeName, &projectName, &salary)

		// check errors
		utils.CheckErr(err)

		employees = append(employees, model.Employee{EmployeeID: employeeId, EmployeeName: employeeName, ProjectName: projectName, Salary: salary})
	}
	println(employees)
	var response = model.JsonResponse{Type: "success", Data: employees}

	json.NewEncoder(w).Encode(response)
}

// create employee

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	println("create employee")
	json.NewDecoder(r.Body).Decode(&employee)
	fmt.Println(employee)
	employeeID := employee.EmployeeID
	employeeName := employee.EmployeeName

	if employeeID == 0 || employeeName == "" {
		response := model.JsonResponse{Type: "error", Message: "You are missing employeeId or employeename parameter."}
		json.NewEncoder(w).Encode(response)
	} else {
		err := repository.CreateEmployee(employee, w)

		utils.CheckErr(err)

		response := model.JsonResponse{Type: "success", Message: "The employee record has been inserted successfully!"}
		json.NewEncoder(w).Encode(response)
	}
}

//deleting employee record

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	var response = model.JsonResponse{}

	err := repository.DeleteEmployee(id)

	// check errors
	utils.CheckErr(err)

	response = model.JsonResponse{Type: "success", Message: "The employee record has been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

func DeleteEmployees(w http.ResponseWriter, r *http.Request) {

	err := repository.DeleteEmployees()

	// check errors
	utils.CheckErr(err)

	utils.PrintMessage("All employees have been deleted successfully!")

	var response = model.JsonResponse{Type: "success", Message: "All employees have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

func UpdateEmployees(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	params := mux.Vars(r)

	id := params["id"]
	json.NewDecoder(r.Body).Decode(&employee)
	employeeID := employee.EmployeeID
	employeeName := employee.EmployeeName

	if employeeID == 0 || employeeName == "" {
		response := model.JsonResponse{Type: "error", Message: "You are missing employeeId or employeename parameter."}
		json.NewEncoder(w).Encode(response)
	} else {

		err := repository.UpdateEmployees(w, employee, id)

		// check errors
		utils.CheckErr(err)

		response := model.JsonResponse{Type: "success", Message: "The employee record has been inserted successfully!"}
		json.NewEncoder(w).Encode(response)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	var employee model.Employee

	json.NewDecoder(r.Body).Decode(&employee)
	employeeID := employee.EmployeeID

	if employeeID == 0 {
		response := model.JsonResponse{Type: "error", Message: "You are missing employeeId parameter."}
		json.NewEncoder(w).Encode(response)
	} else {
		db := utils.SetupDB()

		utils.PrintMessage("get employee details from DB")

		_, err := db.Exec("SELECT * FROM employee.Employee where employeeid = $1", employeeID)

		// check errors
		utils.CheckErr(err)

		//response = JsonResponse{Type: "success", Message: "The employee record has been deleted successfully!"}
		json.NewEncoder(w).Encode(employee)
	}

}
