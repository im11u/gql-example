package domain

// 部署
type Department struct {
	ID   uint
	Name string
}

type DepartmentRepository interface {
	FindAll() ([]*Department, error)
	FindByIDs(ids []uint) (map[uint]*Department, error)
}
