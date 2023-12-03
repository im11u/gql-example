package action

import (
	"database/sql"

	"github.com/im11u/gql-example/go/adapter/repository"
	"github.com/im11u/gql-example/go/domain"
	"github.com/im11u/gql-example/go/usecase"
)

// 従業員に対するアクション
type EmployeeAction interface {
	AllEmployees() ([]*domain.Employee, error)
}

type employeeAction struct {
	findAllUseCase usecase.FindAllEmployeeUseCase
}

func newEmployeeAction(db *sql.DB) *employeeAction {
	repo := repository.NewEmployeeGormRepository(db)

	return &employeeAction{
		findAllUseCase: usecase.NewFindAllEmployeeUseCase(repo),
	}
}

// すべての従業員を取得する
func (a *employeeAction) AllEmployees() ([]*domain.Employee, error) {
	return a.findAllUseCase.Execute()
}
