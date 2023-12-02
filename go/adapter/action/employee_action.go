package action

import (
	"database/sql"

	"github.com/im11u/gql-example/go/adapter/repository"
	"github.com/im11u/gql-example/go/domain"
	"github.com/im11u/gql-example/go/usecase"
)

// 従業員に対するアクション
type EmployeeAction struct {
	findAllUseCase usecase.FindAllEmployeeUseCase
}

func NewEmployeeAction(db *sql.DB) *EmployeeAction {
	repo := repository.NewEmployeeGormRepository(db)

	return &EmployeeAction{
		findAllUseCase: usecase.NewFindAllEmployeeUseCase(repo),
	}
}

// すべての従業員を取得する
func (a *EmployeeAction) FindAll() ([]*domain.Employee, error) {
	return a.findAllUseCase.Execute()
}
