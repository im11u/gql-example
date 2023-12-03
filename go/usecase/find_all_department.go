package usecase

import (
	"github.com/im11u/gql-example/go/domain"
)

// すべての部署を取得する
type FindAllDepartmentUseCase interface {
	Execute() ([]*domain.Department, error)
}

func NewFindAllDepartmentUseCase(repo domain.DepartmentRepository) FindAllDepartmentUseCase {
	return &findAllDepartmentUseCase{
		repo: repo,
	}
}

type findAllDepartmentUseCase struct {
	repo domain.DepartmentRepository
}

func (uc *findAllDepartmentUseCase) Execute() ([]*domain.Department, error) {
	return uc.repo.FindAll()
}
