package dto

import (
	"github.com/Cwilliam/golang-clean-arch/internals/domain"
)

// RouteInput is a struct to represent RouteInput
// It will be used to parse the request (http, cli, etc)
type RouteInput struct {
	Title         string          `json:"title"`
	StartPosition domain.LatLng   `json:"startPosition"`
	EndPosition   domain.LatLng   `json:"endPosition"`
	Points        []domain.LatLng `json:"points"`
}

// RouteOutput is a struct to represent RouteOutput
// It will be used to return the response (http, cli, etc)
type RouteOutput struct {
	Title         string          `json:"title"`
	StartPosition domain.LatLng   `json:"startPosition"`
	EndPosition   domain.LatLng   `json:"endPosition"`
	Points        []domain.LatLng `json:"points"`
}
