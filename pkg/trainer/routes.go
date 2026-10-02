package trainer

import (
	"pokedex/config"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) chi.Router {
	TrainerConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Get("/", TrainerConfigurator.trainersHandler)
	router.Get("/{id}", TrainerConfigurator.trainersByIdHandler)
	return router
}