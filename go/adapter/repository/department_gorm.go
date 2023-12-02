package repository

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/im11u/gql-example/go/domain"
)

// GORMを使用する部署リポジトリ
type DepartmentGormRepository struct {
	db *gorm.DB
}

func NewDepartmentGormRepository(db *sql.DB) *DepartmentGormRepository {
	return &DepartmentGormRepository{
		db: openGorm(db),
	}
}

func (r *DepartmentGormRepository) FindAll() ([]*domain.Department, error) {
	var departments []*domain.Department

	r.db.Order("id").Find(&departments)

	return departments, nil
}
