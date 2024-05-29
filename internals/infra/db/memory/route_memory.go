package infra

import "github.com/Cwilliam/golang-clean-arch/internals/domain"

// It simulates a dabatase in memory
// It will implement RouteRepositoryInterface
// It will have a slice of Route
type RouteInMemoryRepository struct {
	Routes []domain.Route
}

// NewRouteInMemoryRepository is a function to create RouteInMemoryRepository
func NewRouteInMemoryRepository() *RouteInMemoryRepository {
	return &RouteInMemoryRepository{}
}

// Insert is a method to insert a route to the repository
func (r *RouteInMemoryRepository) Insert(route domain.Route) {
	r.Routes = append(r.Routes, route)
}

// List is a method to list all routes in the repository
func (r *RouteInMemoryRepository) List() []domain.Route {
	return r.Routes
}
