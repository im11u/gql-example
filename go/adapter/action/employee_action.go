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
	AllEmployees(ctx context.Context) ([]*domain.Employee, error)
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
func (a *employeeAction) AllEmployees(ctx context.Context) ([]*domain.Employee, error) {
	return a.findAllUseCase.Execute()
}

// 指定の部署の従業員を遅延ロードする
func (a *employeeAction) LoadEmployees(ctx context.Context, departmentID uint) ([]*domain.Employee, error) {
	thunk := a.loader.Load(ctx, departmentID)
	return thunk()
}

// データローダを初期化する
func (a *employeeAction) initLoader() {
	batchFn := func(ctx context.Context, departmentIDs []uint) []*dataloader.Result[[]*domain.Employee] {
		employees, err := a.findUseCase.Execute(departmentIDs)
		return makeLoaderResults(employees, err, len(departmentIDs))
	}

	a.loader = dataloader.NewBatchedLoader(batchFn)
}
