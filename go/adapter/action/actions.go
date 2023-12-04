package action

import (
	"database/sql"

	"github.com/graph-gophers/dataloader/v7"
)

type Actions interface {
	DepartmentAction
	EmployeeAction
}

type actions struct {
	*departmentAction
	*employeeAction
}

func NewActions(db *sql.DB) Actions {
	return &actions{
		departmentAction: newDepartmentAction(db),
		employeeAction:   newEmployeeAction(db),
	}
}

// データローダの結果を生成する
func makeLoaderResults[V any](values []V, err error, size int) []*dataloader.Result[V] {
	results := make([]*dataloader.Result[V], size)

	for i := range results {
		r := &dataloader.Result[V]{}

		if err == nil {
			r.Data = values[i]
		} else {
			r.Error = err
		}

		results[i] = r
	}

	return results
}
