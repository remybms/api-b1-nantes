package models

type Team struct {
	Id       int      `json:"id"`
	Pokemons []string `json:"pokemons"`
}
