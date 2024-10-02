package infra

import "github.com/dinho-carvalho/stock/internal/usecase"

type (
	gatewayImpl struct {
		usecase.Repository
	}
)

func NewGateway(gateway usecase.Repository) gatewayImpl {
	return gatewayImpl{
		Repository: gateway,
	}
}

func (impl *gatewayImpl) Find(name string) ([]string, error) {

	return []string{"BRTW01", "ARBA01"}, nil
}
