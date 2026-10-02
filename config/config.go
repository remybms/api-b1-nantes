package config

import (
	"os"
	"pokedex/database/dbmodels"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	PokemonsRepository dbmodels.PokemonsRepository
	TrainersRepository dbmodels.TrainersRepository
	TeamsRepository    dbmodels.TeamsRepository
}

func New() (*Config, error) {
	godotenv.Load()
	config := Config{}
	dsn := os.Getenv("SUPABASE_DB_URL")
	databaseSession, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	config.PokemonsRepository = dbmodels.NewPokemonRepository(databaseSession)
	config.TrainersRepository = dbmodels.NewTrainersRepository(databaseSession)
	config.TeamsRepository = dbmodels.NewTeamsRepository(databaseSession)
	return &config, nil
}
