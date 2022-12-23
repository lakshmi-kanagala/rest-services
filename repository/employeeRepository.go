package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rest-services/model"
	"rest-services/utils"
)

func GetEmployees() (rows *sql.Rows) {
	utils.PrintMessage("Getting employees...")
	db := utils.SetupDB()

	// Get all employees from employee table
	rows, err := db.Query("SELECT * FROM employee.Employee")

	// check errors
	utils.CheckErr(err)
	return rows
}

func CreateEmployee(employee model.Employee, w http.ResponseWriter) error {
	db := utils.SetupDB()
	json.NewEncoder(w).Encode(&employee)

	utils.PrintMessage("Inserting employee into DB")

	fmt.Println("Inserting new employee with ID: ", employee.EmployeeID, " and name: ", employee.EmployeeName)

	_, err := db.Query("INSERT INTO employee.Employee(employeeid, employeeName, projectName, salary) VALUES($1, $2, $3, $4);", employee.EmployeeID, employee.EmployeeName, employee.ProjectName, employee.Salary)

	// check errors
	utils.CheckErr(err)
	return err
}

func DeleteEmployee(id string) error {
	db := utils.SetupDB()

	utils.PrintMessage("Deleting employee from DB")

	_, err := db.Exec("DELETE FROM employee.Employee where employeeid = $1", id)
	return err
}

func DeleteEmployees() error {
	db := utils.SetupDB()

	utils.PrintMessage("Deleting all employees...")

	_, err := db.Exec("DELETE FROM employee.Employee")
	return err

}

func UpdateEmployees(w http.ResponseWriter, employee model.Employee, id string) error {
	db := utils.SetupDB()
	json.NewEncoder(w).Encode(&employee)

	utils.PrintMessage("updating employee into DB")

	fmt.Println("updating new employee data with ID: ", employee.EmployeeID, " and name: ", employee.EmployeeName)

	_, delerr := db.Exec("DELETE FROM employee.Employee where employeeid = $1", id)

	utils.CheckErr(delerr)

	return CreateEmployee(employee, w)
}
