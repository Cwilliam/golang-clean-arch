package usecases

import (
	"github.com/Cwilliam/golang-clean-arch/internals/domain"
	"github.com/Cwilliam/golang-clean-arch/internals/dto"
)

// RouteUseCase is a struct to represent RouteUseCase
type RouteUseCase struct {
	domain.RouteRepositoryInterface
}

// NewRouteUseCase is a function to create RouteUseCase
func NewRouteUseCase(routeRepo domain.RouteRepositoryInterface) *RouteUseCase {
	return &RouteUseCase{
		RouteRepositoryInterface: routeRepo,
	}
}

// Execute is a function to execute RouteUseCase
// It will insert a new route to the repository
// and return the output
func (uc *RouteUseCase) Execute(input dto.RouteInput) dto.RouteOutput {
	route := domain.Route{
		Title:         input.Title,
		StartPosition: input.StartPosition,
		EndPosition:   input.EndPosition,
		Points:        input.Points,
	}

	uc.RouteRepositoryInterface.Insert(route)

	routeOutput := dto.RouteOutput(route)

	return routeOutput
}
