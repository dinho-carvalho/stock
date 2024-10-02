package usecase

import (
	"context"
)

type (
	Repository interface {
		Find(string) ([]string, error)
	}
	useCase struct {
		Repository
	}
	FindItemUseCase interface {
		Execute(context.Context, string) ([]string, error)
	}
)

func NewUseCase() FindItemUseCase {
	return &useCase{}
}

func (uc *useCase) Execute(ctx context.Context, name string) ([]string, error) {
	return uc.Repository.Find(name)
}
