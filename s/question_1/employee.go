package q1

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

// Reference: https://mari-azevedo.medium.com/s-o-l-i-d-principles-what-are-they-and-why-projects-should-use-them-50b85e4aa8b6
// Employee class/struct case

// Question:
// 1. Is below code conform SRP?
// 2. If not, what part is violating SRP?
// 3. please build a solution that conform SRP

// Some context:
// let employee has SeniorEngineer & JuniorEngineer
// where tax calculation is different

// Connection handle database connection
type Connection struct {
	// DB constring
}

func (c *Connection) Exec(sql string) error {
	// Logic to execute sql query like insert/update
	if len(sql) < 1 {
		return errors.New("fail to execute query")
	}
	fmt.Println("Running query:" + sql)
	return nil
}

type Employee struct {
	id         int
	name       string
	salary     float64
	connection Connection
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) GetSalary() float64 {
	return e.salary
}

func (e *Employee) CalculateSalary() float64 {
	return e.salary - (e.salary * 0.225)
}

// Save employee to database
func (e *Employee) Save() error {
	err := e.connection.Exec("Insert employee " + e.name)
	if err != nil {
		log.Err(err)
		return err
	}
	return nil
}
