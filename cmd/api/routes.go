package api

import (
	"net/http"

	"github.com/smarulanda97/assessment-golang/cmd/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// BindRoutesApi
func BindRoutesApi(mux *chi.Mux) http.Handler {
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Route("/api", func(r chi.Router) {
		r.Post("/quizzes", handlers.QuizStoreHandler)
		r.Get("/questions", handlers.QuestionIndexHandler)
	})

	return mux
}
