package main

import (
	. "github.com/ryeoman/solid/case_employee/business"
)

func main() {
	employeeList := EmployeeList{}
	employees := employeeList.All()

	for _, employee := range employees {
		employee.PrintPaySlip()
	}
}
