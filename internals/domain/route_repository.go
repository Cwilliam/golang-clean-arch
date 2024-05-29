package domain

// RouteRepositoryInterface is an interface to represent RouteRepository
type RouteRepositoryInterface interface {
	Insert(route Route)
	List() []Route
}
