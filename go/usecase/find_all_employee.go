package usecase

import (
	"github.com/im11u/gql-example/go/domain"
)

// すべての従業員を取得する
type FindAllEmployeeUseCase interface {
	Execute() ([]*domain.Employee, error)
}

func NewFindAllEmployeeUseCase(repo domain.EmployeeRepository) FindAllEmployeeUseCase {
	return &findAllEmployeeUseCase{
		repo: repo,
	}
}

type findAllEmployeeUseCase struct {
	repo domain.EmployeeRepository
}

func (uc *findAllEmployeeUseCase) Execute() ([]*domain.Employee, error) {
	return uc.repo.FindAll()
}
