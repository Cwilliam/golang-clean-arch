package usecases

import (
	"github.com/Cwilliam/golang-clean-arch/internals/domain"
	"github.com/Cwilliam/golang-clean-arch/internals/dto"
)

// RouteUseCase is a struct to represent RouteUseCase
type RouteListUseCase struct {
	domain.RouteRepositoryInterface
}

// NewRouteUseCase is a function to create RouteUseCase
func NewRouteListUseCase(routeRepo domain.RouteRepositoryInterface) *RouteListUseCase {
	return &RouteListUseCase{
		RouteRepositoryInterface: routeRepo,
	}
}

// Execute is a function to execute RouteUseCase
// It will insert a new route to the repository
// and return the output
func (uc *RouteListUseCase) Execute() []dto.RouteOutput {

	routes := uc.RouteRepositoryInterface.List()

	routeOutput := []dto.RouteOutput{}
	for _, route := range routes {
		routeOutput = append(routeOutput, dto.RouteOutput(route))
	}

	return routeOutput
}
