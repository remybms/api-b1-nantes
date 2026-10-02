package team

import (
	"net/http"
	"pokedex/config"
	"pokedex/database/dbmodels"
	"pokedex/pkg/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TeamConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *TeamConfigurator {
	return &TeamConfigurator{configuration}
}

func teamToModel(teams []*dbmodels.Team) []models.Team {
	teamsToModel := &models.Team{}
	teamsEdited := []models.Team{}
	for _, team := range teams{
		teamsToModel.Pokemons = team.Pokemons
		teamsEdited = append(teamsEdited, *teamsToModel)
	}
	return teamsEdited
}

func (config *TeamConfigurator) teamsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := config.TeamsRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve teams !"})
	}
	teamsEdited := teamToModel(teams)
	render.JSON(w, r, teamsEdited)
}

func (config *TeamConfigurator) teamByIdHandler(w http.ResponseWriter, r *http.Request) {
	teamId := chi.URLParam(r, "id")
	team, err := config.TeamsRepository.FindById(teamId)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to find this team !"})
	}
	teamEdited := teamToModel(team)
	render.JSON(w, r, teamEdited)
}