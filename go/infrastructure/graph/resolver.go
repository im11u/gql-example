package graph

import (
	"database/sql"

	"github.com/im11u/gql-example/go/adapter/action"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	departmentAction *action.DepartmentAction
	employeeAction   *action.EmployeeAction
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		departmentAction: action.NewDepartmentAction(db),
		employeeAction:   action.NewEmployeeAction(db),
	}
}
