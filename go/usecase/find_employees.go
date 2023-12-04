package usecase

import (
	"github.com/im11u/gql-example/go/domain"
)

// 指定の部署の従業員を取得する
type FindEmployeesUseCase interface {
	Execute(departmentIDs []uint) ([][]*domain.Employee, error)
}

func NewFindEmployeesUseCase(repo domain.EmployeeRepository) FindEmployeesUseCase {
	return &findEmployeesUseCase{
		repo: repo,
	}
}

type findEmployeesUseCase struct {
	repo domain.EmployeeRepository
}

func (uc *findEmployeesUseCase) Execute(departmentIDs []uint) ([][]*domain.Employee, error) {
	employees := make([][]*domain.Employee, len(departmentIDs))

	m, err := uc.repo.FindByDepartmentIDs(departmentIDs)
	if err != nil {
		return nil, err
	}

	// 部署IDと同じ順序で返す
	for i, id := range departmentIDs {
		employees[i] = m[id]
	}

	return employees, nil
}
