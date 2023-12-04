package action

import (
	"context"
	"database/sql"

	"github.com/graph-gophers/dataloader/v7"

	"github.com/im11u/gql-example/go/adapter/repository"
	"github.com/im11u/gql-example/go/domain"
	"github.com/im11u/gql-example/go/usecase"
)

// 部署に対するアクション
type DepartmentAction interface {
	AllDepartments() ([]*domain.Department, error)
	LoadDepartment(ctx context.Context, id uint) (*domain.Department, error)
}

type departmentAction struct {
	findAllUseCase usecase.FindAllDepartmentUseCase
	findUseCase    usecase.FindDepartmentUseCase
	loader         dataloader.Interface[uint, *domain.Department]
}

func newDepartmentAction(db *sql.DB) *departmentAction {
	repo := repository.NewDepartmentGormRepository(db)

	act := &departmentAction{
		findAllUseCase: usecase.NewFindAllDepartmentUseCase(repo),
		findUseCase:    usecase.NewFindDepartmentUseCase(repo),
	}
	act.initLoader()

	return act
}

// すべての部署を取得する
func (a *departmentAction) AllDepartments() ([]*domain.Department, error) {
	return a.findAllUseCase.Execute()
}

// 指定のIDの部署を遅延ロードする
func (a *departmentAction) LoadDepartment(ctx context.Context, id uint) (*domain.Department, error) {
	thunk := a.loader.Load(ctx, id)
	return thunk()
}

func (a *departmentAction) initLoader() {
	batchFn := func(ctx context.Context, ids []uint) []*dataloader.Result[*domain.Department] {
		var results []*dataloader.Result[*domain.Department]
		departments, _ := a.findUseCase.Execute(ids)
		for _, data := range departments {
			results = append(results, &dataloader.Result[*domain.Department]{
				Data: data,
			})
		}
		return results
	}

	a.loader = dataloader.NewBatchedLoader(batchFn)
}
