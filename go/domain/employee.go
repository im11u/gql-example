package domain

// 従業員
type Employee struct {
	ID           uint
	Name         string
	DepartmentID uint
}

type EmployeeRepository interface {
	FindAll() ([]*Employee, error)
	FindByIDs(ids []uint) (map[uint]*Employee, error)
}
