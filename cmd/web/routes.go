package web

import (
	"net/http"

	"github.com/smarulanda97/assesment-golang/cmd/web/handlers"

	"github.com/go-chi/chi/v5"
)

// BindRoutesApi
func BindRoutesWeb(mux *chi.Mux) http.Handler {
	mux.Get("/", handlers.HomeHandler)
	mux.Get("/home", handlers.HomeHandler)

	return mux
}
