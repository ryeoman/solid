package main

import (
	. "github.com/ryeoman/solid/single_responsibility_principle/business"
)

func main() {
	employeeList := EmployeeList{}
	employees := employeeList.All()

	for _, employee := range employees {
		employee.PrintPaySlip()
	}
}
