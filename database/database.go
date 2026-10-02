package database

import (
	"pokedex/database/dbmodels"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodels.Pokemon{},
		&dbmodels.Trainer{},
		&dbmodels.Team{},
	)
	log.Println("Database migrated succesfully")
}