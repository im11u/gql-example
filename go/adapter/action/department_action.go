package action

import (
	"database/sql"

	"github.com/im11u/gql-example/go/adapter/repository"
	"github.com/im11u/gql-example/go/domain"
	"github.com/im11u/gql-example/go/usecase"
)

// 部署に対するアクション
type DepartmentAction interface {
	AllDepartments() ([]*domain.Department, error)
}

type departmentAction struct {
	findAllUseCase usecase.FindAllDepartmentUseCase
}

func newDepartmentAction(db *sql.DB) *departmentAction {
	repo := repository.NewDepartmentGormRepository(db)

	return &departmentAction{
		findAllUseCase: usecase.NewFindAllDepartmentUseCase(repo),
	}
}

// すべての部署を取得する
func (a *departmentAction) AllDepartments() ([]*domain.Department, error) {
	return a.findAllUseCase.Execute()
}
