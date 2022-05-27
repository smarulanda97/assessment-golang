package database

import (
	"errors"

	"github.com/smarulanda97/assessment-golang/internal/models"
)

type MemoryRepository struct {
}

func NewMemoryRepository() (*MemoryRepository, error) {
	return &MemoryRepository{}, nil
}

func (mr *MemoryRepository) GetQuizzers() ([]models.Quizzer, error) {
	quizzer1 := models.Quizzer{
		ID:    1,
		Score: 75,
	}

	quizzer2 := models.Quizzer{
		ID:    2,
		Score: 50,
	}

	quizzer3 := models.Quizzer{
		ID:    3,
		Score: 50,
	}

	quizzer4 := models.Quizzer{
		ID:    4,
		Score: 50,
	}

	quizzer5 := models.Quizzer{
		ID:    5,
		Score: 25,
	}

	quizzer6 := models.Quizzer{
		ID:    6,
		Score: 100,
	}

	return []models.Quizzer{quizzer1, quizzer2, quizzer3, quizzer4, quizzer5, quizzer6}, nil
}

func (mr *MemoryRepository) GetQuestionByID(id int) (models.Question, error) {
	var question models.Question

	questions, err := mr.GetQuestions()
	if err != nil {
		return question, err
	}

	for i := range questions {
		if questions[i].ID == int64(id) {
			return questions[i], nil
		}
	}

	return question, errors.New("")
}

func (mr *MemoryRepository) GetQuestions() ([]models.Question, error) {
	question1 := models.Question{
		ID:            1,
		Description:   "What are SOLID principles?",
		CorrectAnswer: "B",
		Answers: map[string]string{
			"A": "SOLID is the programming language used in the Linux kernel",
			"B": "SOLID makes reference to five principle of software designs more understandable, flexible, and maintainable.",
			"C": "SOLID makes reference to naming convention for variables, functions and classes.",
		},
	}

	question2 := models.Question{
		ID:            2,
		Description:   "What is HTML?",
		CorrectAnswer: "A",
		Answers: map[string]string{
			"A": "HyperText Markup Language",
			"B": "HyperText Model Language",
			"C": "HyperText Model Line",
		},
	}

	question3 := models.Question{
		ID:            3,
		Description:   "What is the most used programming language for websites?",
		CorrectAnswer: "C",
		Answers: map[string]string{
			"A": "Rust",
			"B": "Ruby",
			"C": "JavaScript",
		},
	}

	question4 := models.Question{
		ID:            4,
		Description:   "What is a queue in programming?",
		CorrectAnswer: "A",
		Answers: map[string]string{
			"A": "It is a data structure in which the first element added to the queue will be the first one to be removed.",
			"B": "It is a data structure in which the first element added to the queue will be the last one to be removed",
			"C": "Does Not Know/Does Not Respond",
		},
	}

	return []models.Question{question1, question2, question3, question4}, nil
}
