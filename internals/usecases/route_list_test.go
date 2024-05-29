package usecases

import (
	"testing"

	"github.com/Cwilliam/golang-clean-arch/internals/domain"
	infra "github.com/Cwilliam/golang-clean-arch/internals/infra/db/memory"
)

func TestRouteListUseCase(t *testing.T) {
	routeRepo := infra.NewRouteInMemoryRepository()
	routeUseCase := NewRouteListUseCase(routeRepo)

	routes := routeUseCase.Execute()

	if len(routes) != 0 {
		t.Errorf("Expected 0, got %d", len(routes))
	}
}

func TestRouteListUseCaseWithRoutes(t *testing.T) {
	routeRepo := infra.NewRouteInMemoryRepository()
	routeUseCase := NewRouteListUseCase(routeRepo)

	routeRepo.Insert(domain.Route{
		Title:         "test",
		StartPosition: domain.LatLng{Lat: 1.0, Lng: 2.0},
		EndPosition:   domain.LatLng{Lat: 3.0, Lng: 4.0},
		Points:        []domain.LatLng{{Lat: 5.0, Lng: 6.0}, {Lat: 7.0, Lng: 8.0}},
	})

	routes := routeUseCase.Execute()

	if len(routes) != 1 {
		t.Errorf("Expected 1, got %d", len(routes))
	}

	if routes[0].Title != "test" {
		t.Errorf("Expected test, got %s", routes[0].Title)
	}
}
