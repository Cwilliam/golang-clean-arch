package usecases

import (
	"testing"

	"github.com/Cwilliam/golang-clean-arch/internals/domain"
	"github.com/Cwilliam/golang-clean-arch/internals/dto"
	infra "github.com/Cwilliam/golang-clean-arch/internals/infra/db/memory"
)

func TestRouteUseCase(t *testing.T) {
	repositoryInMemory := infra.NewRouteInMemoryRepository()
	useCase := NewRouteUseCase(repositoryInMemory)

	input := dto.RouteInput{
		Title:         "test",
		StartPosition: domain.LatLng{Lat: 1.0, Lng: 2.0},
		EndPosition:   domain.LatLng{Lat: 3.0, Lng: 4.0},
		Points:        []domain.LatLng{{Lat: 5.0, Lng: 6.0}, {Lat: 7.0, Lng: 8.0}},
	}

	output := useCase.Execute(input)

	if len(repositoryInMemory.Routes) != 1 {
		t.Errorf("Expected 1 route, got %d", len(repositoryInMemory.Routes))
	}

	if !compareLatLng(repositoryInMemory.Routes[0].StartPosition, input.StartPosition) ||
		!compareLatLng(repositoryInMemory.Routes[0].EndPosition, input.EndPosition) ||
		!compareLatLngSlice(repositoryInMemory.Routes[0].Points, input.Points) {
		t.Errorf("Expected route %v, got %v", input, repositoryInMemory.Routes[0])
	}

	if output.Title != input.Title {
		t.Errorf("Expected title %s, got %s", input.Title, output.Title)
	}
}

// Custom comparison function for domain.LatLng
func compareLatLng(a, b domain.LatLng) bool {
	return a.Lat == b.Lat && a.Lng == b.Lng
}

func compareLatLngSlice(a, b []domain.LatLng) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !compareLatLng(a[i], b[i]) {
			return false
		}
	}
	return true
}
