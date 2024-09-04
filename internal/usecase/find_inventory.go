package usecase

import (
	"context"
)
type ( 

	useCase struct {

}
 FindInventoryUseCase interface {
	Execute(context.Context, string) ([]string, error)
 } 
)

func NewUseCase() FindInventoryUseCase {
	return &useCase{}
}

func (uc *useCase) Execute(ctx context.Context name string) ([]string, error) {
	return nil
}

