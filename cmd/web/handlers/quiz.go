package handlers

import (
	"net/http"

	"github.com/smarulanda97/assesment-golang/cmd/web/render"
)

type QuizResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func QuizHandler(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderTemplate(w, r, "quiz", &render.TemplateData{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
