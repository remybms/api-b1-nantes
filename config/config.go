package config

import (
	"os"
	"pokedex/database"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
}

func New() (*Config, error) {
	godotenv.Load()
	config := Config{}
	dsn := os.Getenv("SUPABASE_DB_URL")
	databaseSession, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)
	return &config, nil
}
