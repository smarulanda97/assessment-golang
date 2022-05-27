package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/smarulanda97/assessment-golang/cmd/web/handlers"
)

// BindRoutesApi
func BindRoutesWeb(mux *chi.Mux) http.Handler {
	mux.Get("/", handlers.QuizHandler)
	mux.Get("/quiz", handlers.QuizHandler)

	return mux
}
