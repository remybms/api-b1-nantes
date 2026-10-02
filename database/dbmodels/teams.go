package dbmodels

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Pokemons pq.StringArray `gorm:"type:text[]"`
}

type TeamsRepository interface {
	FindAll()([]*Team, error)
	FindById(teamId int)([]*Team, error)
}

type teamsRepository struct {
	db *gorm.DB
}

func NewTeamsRepository(db *gorm.DB) TeamsRepository{
	return &teamsRepository{db: db}
}

func (r *teamsRepository) FindAll()([]*Team, error){
	var Teams []*Team
	if err := r.db.Find(&Teams).Error; err != nil{
		return nil, err
	}
	return Teams, nil
}

func (r *teamsRepository) FindById(teamId int) ([]*Team, error){
	var entry []*Team
	if err := r.db.Where("id = ?", teamId).Find(&entry).Error; err != nil{
		return nil, err
	}
	return entry, nil
}