package usecase

import (
	"github.com/im11u/gql-example/go/domain"
)

// 指定のIDの部署を取得する
type FindDepartmentUseCase interface {
	Execute(ids []uint) ([]*domain.Department, error)
}

func NewFindDepartmentUseCase(repo domain.DepartmentRepository) FindDepartmentUseCase {
	return &findDepartmentUseCase{
		repo: repo,
	}
}

type findDepartmentUseCase struct {
	repo domain.DepartmentRepository
}

func (uc *findDepartmentUseCase) Execute(ids []uint) ([]*domain.Department, error) {
	departments := make([]*domain.Department, len(ids))

	m, err := uc.repo.FindByIDs(ids)
	if err != nil {
		return nil, err
	}

	// IDと同じ順序で返す
	for i, id := range ids {
		departments[i] = m[id]
	}

	return departments, nil
}
