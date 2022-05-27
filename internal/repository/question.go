package repository

import "github.com/smarulanda97/assessment-golang/internal/models"

type QuestionRepository interface {
	GetQuestions() ([]models.Question, error)
	GetQuestionByID(id int) (models.Question, error)
}

var implementation QuestionRepository

func SetRepository(repository QuestionRepository) {
	implementation = repository
}

func GetQuestions() ([]models.Question, error) {
	return implementation.GetQuestions()
}

func GetQuestionByID(id int) (models.Question, error) {
	return implementation.GetQuestionByID(id)
}
