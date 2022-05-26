package repository

import "github.com/smarulanda97/assesment-golang/internal/models"

type QuestionRepository interface {
	GetQuestions() ([]models.Question, error)
}

var implementation QuestionRepository

func SetRepository(repository QuestionRepository) {
	implementation = repository
}

func GetQuestions() ([]models.Question, error) {
	return implementation.GetQuestions()
}
