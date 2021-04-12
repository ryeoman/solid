package business

import (
	. "github.com/ryeoman/solid/case_employee/entity"
)

// define employee position
var (
	SeniorDeveloper = NewPosition(&EleventPercentRule{})
	JuniorDeveloper = NewPosition(&EleventPercentRule{})
)

type TwentyTwoAndHalfPercentRule struct {
}

func (tth *TwentyTwoAndHalfPercentRule) Calculates(employee Employee) float64 {
	return employee.Salary - (employee.Salary * 0.225)
}

type EleventPercentRule struct {
}

func (ep *EleventPercentRule) Calculates(employee Employee) float64 {
	return employee.Salary - (employee.Salary * 0.11)
}
