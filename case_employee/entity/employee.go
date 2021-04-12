package entity

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID       int
	Name     string
	Salary   float64
	Position Position
}

func NewEmployee(id int, name string, salary float64, position Position) *Employee {
	return &Employee{
		ID:       id,
		Name:     name,
		Salary:   salary,
		Position: position,
	}
}

// CalculateSalary after tax adjustment
func (e *Employee) CalculateSalary() float64 {
	return e.Position.rule.Calculates(*e)
}

func (e *Employee) PrintPaySlip() {
	payslip := map[string]interface{}{
		"name":             e.Name,
		"salary after tax": fmt.Sprintf("Rp.%.2f", e.CalculateSalary()),
	}
	b, err := json.MarshalIndent(payslip, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
