package domain

// 従業員
type Employee struct {
	ID   uint
	Name string
}

type EmployeeRepository interface {
	FindAll() ([]*Employee, error)
}
