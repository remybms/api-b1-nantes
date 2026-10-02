package pokemon

import (
	"net/http"
	"pokedex/config"
	"pokedex/database/dbmodels"
	"pokedex/pkg/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type PokemonConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *PokemonConfigurator {
	return &PokemonConfigurator{configuration}
}

func pokemonsToModel(pokemons []*dbmodels.Pokemon) []models.Pokemon {
	pokemonToModel := &models.Pokemon{}
	pokemonsEdited := []models.Pokemon{}
	for _, pokemon := range pokemons {
		pokemonToModel.Name = pokemon.Name
		pokemonToModel.Image = pokemon.Image
		pokemonToModel.Description = pokemon.Description
		pokemonToModel.FirstType = pokemon.FirstType
		pokemonToModel.SecondType = pokemon.SecondType
		pokemonToModel.Generation = pokemon.Generation
		pokemonToModel.SignatureAttack = pokemon.SignatureAttack
		pokemonsEdited = append(pokemonsEdited, *pokemonToModel)
	}
	return pokemonsEdited
}

func (config *PokemonConfigurator) pokemonsHandler(w http.ResponseWriter, r *http.Request) {
	pokemons, err := config.PokemonsRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the Pokemons !"})
		return
	}
	pokemonsEdited := pokemonsToModel(pokemons)
	render.JSON(w, r, pokemonsEdited)
}

func (config *PokemonConfigurator) pokemonByIdHandler(w http.ResponseWriter, r *http.Request) {
	pokemonId := chi.URLParam(r, "id")
	pokemon, err := config.PokemonsRepository.FindById(pokemonId)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error" : "Failed to load the Pokemon"})
	}
	pokemonEdited := pokemonsToModel(pokemon)
	render.JSON(w, r, pokemonEdited)
}