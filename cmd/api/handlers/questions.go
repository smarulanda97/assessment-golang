package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smarulanda97/assesment-golang/internal/models"
	"github.com/smarulanda97/assesment-golang/internal/repository"
)

type QuestionsResponse struct {
	Message   string            `json:"message"`
	Questions []models.Question `json:"questions"`
	Status    bool              `json:"status"`
}

// QuestionsIndexHandler
func QuestionsIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	questions, err := repository.GetQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(QuestionsResponse{
		Message:   "OK",
		Status:    true,
		Questions: questions,
	})
}
