package domain

// LatLng represents a latitude and longitude
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Route represents a route with a title, start and end position and a list of points
type Route struct {
	Title         string   `json:"title"`
	StartPosition LatLng   `json:"startPosition"`
	EndPosition   LatLng   `json:"endPosition"`
	Points        []LatLng `json:"points"`
}

// NewRoute creates a new route with the given title, start and end position and points
func NewRoute(title string, startPosition LatLng, endPosition LatLng, points []LatLng) *Route {
	return &Route{
		Title:         title,
		StartPosition: startPosition,
		EndPosition:   endPosition,
		Points:        points,
	}
}

// UpdateTitle updates the title of the route
func (r *Route) UpdateTitle(title string) {
	r.Title = title
}

// UpdatePosition updates the start and end position of the route
func (r *Route) UpdatePosition(startPosition LatLng, endPosition LatLng) {
	r.StartPosition = startPosition
	r.EndPosition = endPosition
}

// UpdatePoints updates the points of the route
func (r *Route) UpdatePoints(points []LatLng) {
	r.Points = points
}
