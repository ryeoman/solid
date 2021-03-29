package business

import (
	. "github.com/ryeoman/solid/single_responsibility_principle/dao"
	. "github.com/ryeoman/solid/single_responsibility_principle/entity"
)

type EmployeeList struct {
	employeeDAO EmployeeDAO
}

func (el *EmployeeList) List() []Employee {
	employeeRaws := el.employeeDAO.List()

	employees := make([]Employee, 0, len(employeeRaws))
	for _, employeeRaw := range employeeRaws {
		employees = append(employees, Employee{
			ID:       employeeRaw.ID,
			Name:     employeeRaw.Name,
			Salary:   employeeRaw.Salary,
			Position: el.mapPosition(employeeRaw.Position),
		})
	}
	return employees
}

func (el *EmployeeList) mapPosition(positionRaw int) Position {
	var position Position
	switch positionRaw {
	case 1:
		position = *JuniorDeveloper
	case 2:
		position = *SeniorDeveloper
	}
	return position
}
