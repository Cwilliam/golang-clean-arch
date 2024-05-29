package domain

import (
	"reflect"
	"testing"
)

func TestCreateRoute(t *testing.T) {
	start := LatLng{Lat: 1.0, Lng: 2.0}
	end := LatLng{Lat: 3.0, Lng: 4.0}
	points := []LatLng{{Lat: 5.0, Lng: 6.0}, {Lat: 7.0, Lng: 8.0}}
	route := NewRoute("test", start, end, points)

	if route.Title != "test" {
		t.Errorf("expected title to be 'test', got %s", route.Title)
	}

	if route.StartPosition != start {
		t.Errorf("expected start position to be %v, got %v", start, route.StartPosition)
	}

	if route.EndPosition != end {
		t.Errorf("expected end position to be %v, got %v", end, route.EndPosition)
	}

	if !reflect.DeepEqual(route.Points, points) {
		t.Errorf("expected points to be %v, got %v", points, route.Points)
	}
}

func TestCreateRouteWithPointsAsNil(t *testing.T) {
	start := LatLng{Lat: 1.0, Lng: 2.0}
	end := LatLng{Lat: 3.0, Lng: 4.0}
	route := NewRoute("test", start, end, nil)

	if route.Title != "test" {
		t.Errorf("expected title to be 'test', got %s", route.Title)
	}

	if route.StartPosition != start {
		t.Errorf("expected start position to be %v, got %v", start, route.StartPosition)
	}

	if route.EndPosition != end {
		t.Errorf("expected end position to be %v, got %v", end, route.EndPosition)
	}

	if route.Points != nil {
		t.Errorf("expected points to be nil, got %v", route.Points)
	}
}

func TestUpdateTitle(t *testing.T) {
	route := NewRoute("test", LatLng{}, LatLng{}, nil)
	route.UpdateTitle("new title")

	if route.Title != "new title" {
		t.Errorf("expected title to be 'new title', got %s", route.Title)
	}
}

func TestUpdatePosition(t *testing.T) {
	route := NewRoute("test", LatLng{Lat: 1.0, Lng: 2.0}, LatLng{Lat: 3.0, Lng: 4.0}, nil)
	route.UpdatePosition(LatLng{Lat: 5.0, Lng: 6.0}, LatLng{Lat: 7.0, Lng: 8.0})

	startPosition := LatLng{Lat: 5.0, Lng: 6.0}
	if route.StartPosition != startPosition {
		t.Errorf("expected start position to be %v, got %v", startPosition, route.StartPosition)
	}

	endPosition := LatLng{Lat: 7.0, Lng: 8.0}
	if route.EndPosition != endPosition {
		t.Errorf("expected end position to be %v, got %v", endPosition, route.EndPosition)
	}
}

func TestUpdatePoints(t *testing.T) {
	route := NewRoute("test", LatLng{}, LatLng{}, nil)
	points := []LatLng{{Lat: 1.0, Lng: 2.0}, {Lat: 3.0, Lng: 4.0}}
	route.UpdatePoints(points)

	if len(route.Points) != 2 {
		t.Errorf("expected points to have length 2, got %d", len(route.Points))
	}

	if !reflect.DeepEqual(route.Points, points) {
		t.Errorf("expected points to be %v, got %v", points, route.Points)
	}
}
