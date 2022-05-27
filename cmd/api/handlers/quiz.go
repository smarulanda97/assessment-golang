package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smarulanda97/assesment-golang/internal/repository"
)

type QuizzesStorePayload struct {
	Answers map[int]string `json:"answers"`
}

type Result struct {
	TotalQuestions   int     `json:"total_questions"`
	CorrectAnswers   int     `json:"correct_anwsers"`
	InCorrectAnswers int     `json:"incorrect_anwsers"`
	Score            float64 `json:"score"`
}

type QuizzesStoreResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Result  Result `json:"result"`
}

// Process answers submitted by the user and return the result
func QuizStoreHandler(w http.ResponseWriter, r *http.Request) {
	var payload QuizzesStorePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := calculateQuizzesResult(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(QuizzesStoreResponse{
		Message: "OK",
		Status:  true,
		Result:  *result,
	})
}

// Calculate quantity of corrects, incorrect anwers and total of questions
func calculateQuizzesResult(payload *QuizzesStorePayload) (*Result, error) {
	result := Result{}
	totalQuestions := 0
	correctAnswers := 0
	incorrectAnswers := 0

	for questionID, userAnswer := range payload.Answers {
		currentQuestion, err := repository.GetQuestionByID(questionID)
		if err != nil {
			return &result, err
		}

		if currentQuestion.CorrectAnswer == userAnswer {
			correctAnswers++
		} else {
			incorrectAnswers++
		}

		totalQuestions++
	}

	result.TotalQuestions = totalQuestions
	result.CorrectAnswers = correctAnswers
	result.InCorrectAnswers = incorrectAnswers
	result.Score = (float64(correctAnswers) / float64(totalQuestions)) * 100

	return &result, nil
}
