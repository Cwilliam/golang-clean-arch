package main

import (
	"log"
	"net/http"

	infra "github.com/Cwilliam/golang-clean-arch/internals/infra/db/memory"
	httpServer "github.com/Cwilliam/golang-clean-arch/internals/infra/http"
)

func main() {
	repositoryInMemory := infra.NewRouteInMemoryRepository()
	serverMux := httpServer.NewServerHttpSample(repositoryInMemory)

	log.Fatal(http.ListenAndServe(":8080", serverMux.Serve()))
}
