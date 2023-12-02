package usecase

import (
	"github.com/im11u/gql-example/go/domain"
)

// すべての従業員を取得する
type FindAllEmployeeUseCase interface {
	Execute() ([]*domain.Employee, error)
}

func NewFindAllEmployeeUseCase(repo domain.EmployeeRepository) FindAllEmployeeUseCase {
	return &FindAllEmployeeUseCaseImpl{
		repo: repo,
	}
}

type FindAllEmployeeUseCaseImpl struct {
	repo domain.EmployeeRepository
}

func (uc *FindAllEmployeeUseCaseImpl) Execute() ([]*domain.Employee, error) {
	return uc.repo.FindAll()
}
