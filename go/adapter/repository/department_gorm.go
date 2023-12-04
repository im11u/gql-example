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

func (r *DepartmentGormRepository) FindByIDs(ids []uint) (map[uint]*domain.Department, error) {
	m := make(map[uint]*domain.Department)
	var departments []*domain.Department

	r.db.Find(&departments, ids)

	for _, v := range departments {
		m[v.ID] = v
	}

	return m, nil
}
