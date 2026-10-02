package dbmodels

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	OriginRegion string `json:"origin_region"`
	OriginCity   string `json:"origin_city"`
	BestBuddy    string `json:"best_buddy"`
	Team         string `json:"team"`
}


type TrainersRepository interface {
	FindAll()([]*Trainer, error)
	FindById(pokemonId int)([]*Trainer, error)
}

type trainersRepository struct {
	db *gorm.DB
}

func NewTrainersRepository(db *gorm.DB) TrainersRepository{
	return &trainersRepository{db: db}
}

func (r *trainersRepository) FindAll()([]*Trainer, error){
	var trainers []*Trainer
	if err := r.db.Find(&trainers).Error; err != nil{
		return nil, err
	}
	return trainers, nil
}

func (r *trainersRepository) FindById(trainerId int) ([]*Trainer, error){
	var entry []*Trainer
	if err := r.db.Where("id = ?", trainerId).Find(&entry).Error; err != nil{
		return nil, err
	}
	return entry, nil
}