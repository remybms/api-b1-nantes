package trainer

import (
	"net/http"
	"pokedex/config"
	"pokedex/database/dbmodels"
	"pokedex/pkg/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TrainerConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *TrainerConfigurator {
	return &TrainerConfigurator{configuration}
}

func trainerToModel(trainers []*dbmodels.Trainer) []models.Trainer {
	trainerToModel := &models.Trainer{}
	trainersEdited := []models.Trainer{}
	for _, trainer := range trainers {
		trainerToModel.Name = trainer.Name
		trainerToModel.Age = trainer.Age
		trainerToModel.OriginRegion = trainer.OriginRegion
		trainerToModel.OriginCity = trainer.OriginCity
		trainerToModel.BestBuddy = trainer.BestBuddy
		trainerToModel.Team = trainer.Team
		trainersEdited = append(trainersEdited, *trainerToModel)
	}
	return trainersEdited
}

func (config *TrainerConfigurator) trainersHandler(w http.ResponseWriter, r *http.Request) {
	trainers, err := config.TrainersRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the trainers !"})
	}
	trainersEdited := trainerToModel(trainers)
	render.JSON(w, r, trainersEdited)
}

func (config *TrainerConfigurator) trainersByIdHandler(w http.ResponseWriter, r *http.Request){
	trainerId := chi.URLParam(r, "id")
	trainer, err := config.TrainersRepository.FindById(trainerId)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to load the wanted trainer !"})
	}
	trainerEdited := trainerToModel(trainer)
	render.JSON(w, r, trainerEdited)
}