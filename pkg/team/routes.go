package team

import (
	"pokedex/config"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) chi.Router {
	TeamConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Get("/", TeamConfigurator.teamsHandler)
	router.Get("/{id}", TeamConfigurator.teamByIdHandler)
	return router
}