package infra

import (
	"testing"

	"github.com/Cwilliam/golang-clean-arch/internals/domain"
)

func TestRouteInMemoryRepository(t *testing.T) {
	repository := NewRouteInMemoryRepository()

	route := domain.Route{
		Title:         "test",
		StartPosition: domain.LatLng{Lat: 1.0, Lng: 2.0},
		EndPosition:   domain.LatLng{Lat: 3.0, Lng: 4.0},
		Points:        []domain.LatLng{{Lat: 5.0, Lng: 6.0}, {Lat: 7.0, Lng: 8.0}},
	}

	repository.Insert(route)

	if len(repository.Routes) != 1 {
		t.Errorf("Expected 1 route, got %d", len(repository.Routes))
	}

	if !compareLatLng(repository.Routes[0].StartPosition, route.StartPosition) ||
		!compareLatLng(repository.Routes[0].EndPosition, route.EndPosition) ||
		!compareLatLngSlice(repository.Routes[0].Points, route.Points) {
		t.Errorf("Expected route %v, got %v", route, repository.Routes[0])
	}
}

// Custom comparison function for domain.LatLng
func compareLatLng(a, b domain.LatLng) bool {
	return a.Lat == b.Lat && a.Lng == b.Lng
}

// Custom comparison function for []domain.LatLng
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
