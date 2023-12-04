package repository

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/im11u/gql-example/go/domain"
)

// GORMを使用する従業員リポジトリ
type EmployeeGormRepository struct {
	db *gorm.DB
}

func NewEmployeeGormRepository(db *sql.DB) *EmployeeGormRepository {
	return &EmployeeGormRepository{
		db: openGorm(db),
	}
}

func (r *EmployeeGormRepository) FindAll() ([]*domain.Employee, error) {
	var employees []*domain.Employee

	r.db.Order("id").Find(&employees)

	return employees, nil
}

func (r *EmployeeGormRepository) FindByDepartmentIDs(departmentIDs []uint) (map[uint][]*domain.Employee, error) {
	m := make(map[uint][]*domain.Employee)
	var employees []*domain.Employee

	r.db.Where("department_id IN ?", departmentIDs).Order("id").Find(&employees)

	for _, v := range employees {
		m[v.DepartmentID] = append(m[v.DepartmentID], v)
	}

	return m, nil
}
