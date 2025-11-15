package services

import "qa_service/internal/models"

type QuestionService interface {
	Create(text string) (*models.Question, error)
	GetAll() ([]models.Question, error)
	GetByID(id uint) (*models.Question, error)
	Delete(id uint) error
}

type AnswerService interface {
	Create(questionID uint, userID, text string) (*models.Answer, error)
	GetByID(id uint) (*models.Answer, error)
	Delete(id uint) error
}
