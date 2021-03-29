package dao

type EmployeeDAO struct {
	// any data connection
}

type EmployeeRaw struct {
	ID       int
	Name     string
	Salary   float64
	Position int
}

func NewEmployeeDAO() *EmployeeDAO {
	return &EmployeeDAO{}
}

func (ed *EmployeeDAO) List() []EmployeeRaw {
	return []EmployeeRaw{
		{
			ID:       1,
			Name:     "Nico Ivan Cahyadi",
			Salary:   24000000,
			Position: 2,
		},
		{
			ID:       2,
			Name:     "Juan",
			Salary:   20000000,
			Position: 1,
		},
	}
}
