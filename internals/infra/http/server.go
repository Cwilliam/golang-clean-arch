package infra

import (
	"encoding/json"
	"net/http"

	"github.com/Cwilliam/golang-clean-arch/internals/domain"
	"github.com/Cwilliam/golang-clean-arch/internals/dto"
	"github.com/Cwilliam/golang-clean-arch/internals/usecases"
)

// ServerHttpSample is a struct to represent HttpServer
// It will implement Serve method to serve the http server
// It will also have a RouteRepositoryInterface
type ServerHttpSample struct {
	http.Server
	repo domain.RouteRepositoryInterface
}

// NewServerHttpSample is a function to create HttpServer
func NewServerHttpSample(repo domain.RouteRepositoryInterface) *ServerHttpSample {
	return &ServerHttpSample{
		repo: repo,
	}
}

// Serve is a method to serve the http server
func (s *ServerHttpSample) Serve() *http.ServeMux {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("POST /route", s.CreateRoute)
	serverMux.HandleFunc("GET /route/list", s.ListRoute)

	return serverMux
}

// CreateRoute is a method to create a route
// It will parse the request body to RouteInput
// and return the response as RouteOutput
func (s *ServerHttpSample) CreateRoute(w http.ResponseWriter, r *http.Request) {

	routeUseCase := usecases.NewRouteUseCase(s.repo)

	route := dto.RouteInput{}
	err := json.NewDecoder(r.Body).Decode(&route)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	routeOutput := routeUseCase.Execute(route)
	response, err := json.Marshal(routeOutput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// ListRoute is a method to list all routes
// It will return the response as RouteOutput
func (s *ServerHttpSample) ListRoute(w http.ResponseWriter, r *http.Request) {
	routeUseCase := usecases.NewRouteListUseCase(s.repo)

	routes := routeUseCase.Execute()
	response, err := json.Marshal(routes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
