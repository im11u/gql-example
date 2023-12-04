package action

import (
	"context"
	"database/sql"

	"github.com/graph-gophers/dataloader/v7"

	"github.com/im11u/gql-example/go/adapter/repository"
	"github.com/im11u/gql-example/go/domain"
	"github.com/im11u/gql-example/go/usecase"
)

// 従業員に対するアクション
type EmployeeAction interface {
	AllEmployees() ([]*domain.Employee, error)
	LoadEmployees(ctx context.Context, departmentID uint) ([]*domain.Employee, error)
}

type employeeAction struct {
	findAllUseCase usecase.FindAllEmployeeUseCase
	findUseCase    usecase.FindEmployeesUseCase
	loader         dataloader.Interface[uint, []*domain.Employee]
}

func newEmployeeAction(db *sql.DB) *employeeAction {
	repo := repository.NewEmployeeGormRepository(db)

	act := &employeeAction{
		findAllUseCase: usecase.NewFindAllEmployeeUseCase(repo),
		findUseCase:    usecase.NewFindEmployeesUseCase(repo),
	}
	act.initLoader()

	return act
}

// すべての従業員を取得する
func (a *employeeAction) AllEmployees() ([]*domain.Employee, error) {
	return a.findAllUseCase.Execute()
}

// 指定の部署の従業員を遅延ロードする
func (a *employeeAction) LoadEmployees(ctx context.Context, departmentID uint) ([]*domain.Employee, error) {
	thunk := a.loader.Load(ctx, departmentID)
	return thunk()
}

func (a *employeeAction) initLoader() {
	batchFn := func(ctx context.Context, departmentIDs []uint) []*dataloader.Result[[]*domain.Employee] {
		var results []*dataloader.Result[[]*domain.Employee]
		employees, _ := a.findUseCase.Execute(departmentIDs)
		for _, data := range employees {
			results = append(results, &dataloader.Result[[]*domain.Employee]{
				Data: data,
			})
		}
		return results
	}

	a.loader = dataloader.NewBatchedLoader(batchFn)
}
