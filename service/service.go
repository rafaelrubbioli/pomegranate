package service

import (
	"fmt"
	"log"
	"net/http"
	"pomegranate/database"
	"pomegranate/newznab"
	"pomegranate/themoviedb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Tmdb themoviedb.Themoviedb
	Newz []newznab.Newznab
	DB   database.DB
}

type MovieEntry struct {
	Runtime  int32
	Released string
	ImdbId   string
	TmdbId   int32
	Year     int32
	Genres   []string
	Titles   []string
	Images   struct {
		Posters []string
	}
}

type MovieSearchResponse struct {
	Movies []MovieEntry `json:"movies"`
}

func internalError(w http.ResponseWriter, format string, a ...interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("internal error"))
	if err != nil {
		log.Println(fmt.Errorf("http.ResponseWriter.Write: %w", err))
	}
	log.Println(fmt.Errorf(format, a...))
}

func Service(config Config) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pomegranate"))
		if err != nil {
			log.Println(fmt.Errorf("http.ResponseWriter.Write: %w", err))
		}
	})
	r.Get("/movie/search", config.movieSearchHandler)
	r.Get("/movie/add", config.movieAddHandler)

	return r
}
