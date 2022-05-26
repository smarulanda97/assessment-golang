package database

import "github.com/smarulanda97/assesment-golang/internal/models"

type MemoryRepository struct {
}

func NewMemoryRepository() (*MemoryRepository, error) {
	return &MemoryRepository{}, nil
}

func (mr *MemoryRepository) GetQuestions() ([]models.Question, error) {
	question1 := models.Question{
		Id:            1,
		Description:   "What are SOLID principles?",
		CorrectAnswer: "B",
		Answers: map[string]string{
			"A": "SOLID is the programming language used in the Linux kernel",
			"B": "SOLID makes reference to five principle of software designs more understandable, flexible, and maintainable.",
			"C": "SOLID makes reference to naming convention for variables, functions and classes.",
		},
	}

	question2 := models.Question{
		Id:            2,
		Description:   "What is HTML?",
		CorrectAnswer: "A",
		Answers: map[string]string{
			"A": "HyperText Markup Language",
			"B": "HyperText Model Language",
			"C": "HyperText Model Line",
		},
	}

	question3 := models.Question{
		Id:            3,
		Description:   "What is the most used programming language for websites?",
		CorrectAnswer: "C",
		Answers: map[string]string{
			"A": "Rust",
			"B": "Ruby",
			"C": "JavaScript",
		},
	}

	question4 := models.Question{
		Id:            4,
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
