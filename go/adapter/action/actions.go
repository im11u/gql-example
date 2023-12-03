package action

import (
	"database/sql"
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
