package action

import (
	"database/sql"

	"github.com/im11u/gql-example/go/adapter/repository"
	"github.com/im11u/gql-example/go/domain"
	"github.com/im11u/gql-example/go/usecase"
)

// 部署に対するアクション
type DepartmentAction struct {
	findAllUseCase usecase.FindAllDepartmentUseCase
}

func NewDepartmentAction(db *sql.DB) *DepartmentAction {
	repo := repository.NewDepartmentGormRepository(db)

	return &DepartmentAction{
		findAllUseCase: usecase.NewFindAllDepartmentUseCase(repo),
	}
}

// すべての部署を取得する
func (a *DepartmentAction) FindAll() ([]*domain.Department, error) {
	return a.findAllUseCase.Execute()
}
