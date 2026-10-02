package pokemon

import (
	"pokedex/config"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) chi.Router {
	PokemonConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Get("/", PokemonConfigurator.pokemonsHandler)
	router.Get("/{id}", PokemonConfigurator.pokemonByIdHandler)
	return router
}