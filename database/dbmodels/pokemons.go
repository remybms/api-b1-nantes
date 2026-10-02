package dbmodels

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Name            string `json:"name"`
	Image           string `json:"image"`
	Description     string `json:"description"`
	FirstType       string `json:"first_type"`
	SecondType      string `json:"second_type"`
	Generation      int    `json:"generation"`
	SignatureAttack string `json:"signature_attack"`
}

type PokemonsRepository interface {
	FindAll()([]*Pokemon, error)
	FindById(pokemonId int)([]*Pokemon, error)
}

type pokemonsRepository struct {
	db *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) PokemonsRepository{
	return &pokemonsRepository{db: db}
}

func (r *pokemonsRepository) FindAll()([]*Pokemon, error){
	var pokemons []*Pokemon
	if err := r.db.Find(&pokemons).Error; err != nil{
		return nil, err
	}
	return pokemons, nil
}

func (r *pokemonsRepository) FindById(pokemonId int) ([]*Pokemon, error){
	var entry []*Pokemon
	if err := r.db.Where("id = ?", pokemonId).Find(&entry).Error; err != nil{
		return nil, err
	}
	return entry, nil
}