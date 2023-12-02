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
