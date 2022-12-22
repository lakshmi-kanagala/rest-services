package main

type Employee struct {
	EmployeeID   int    `json:"employeeId"`
	EmployeeName string `json:"employeeName"`
	ProjectName  string `json:"projectName"`
	Salary       int    `json:"salary"`
}

type JsonResponse struct {
	Type    string     `json:"type"`
	Data    []Employee `json:"data"`
	Message string     `json:"message"`
}
