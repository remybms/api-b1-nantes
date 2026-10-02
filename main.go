package main

import (
	"log"
	"net/http"
	"pokedex/config"
	"pokedex/docs"
	"pokedex/pkg/pokemon"
	"pokedex/pkg/team"
	"pokedex/pkg/trainer"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter();

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	router.Mount("/api/pokemons", pokemon.Routes(configuration))
	router.Mount("/api/trainers", trainer.Routes(configuration))
	router.Mount("/api/teams", team.Routes(configuration))

	router.Get("/docs/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(docs.SwaggerYAML)
	})

	router.Get("/*", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.yaml"),
	))

	return router
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env :", err)
	}

	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	router := Routes(configuration)

	log.Fatal(http.ListenAndServe(":8080", router))
}
