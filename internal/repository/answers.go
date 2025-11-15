package repository

import (
	"errors"

	"qa_service/internal/models"

	"gorm.io/gorm"
)

type AnswersRepository interface {
	Create(answer *models.Answer) error
	GetByID(id uint) (*models.Answer, error)
	Delete(id uint) error
}

type answersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) AnswersRepository {
	return &answersRepository{db: db}
}

// CREATE with check question exists
func (r *answersRepository) Create(a *models.Answer) error {
	var q models.Question
	// Check if question exists
	if err := r.db.First(&q, a.QuestionID).Error; err != nil {
		return errors.New("question does not exist")
	}

	return r.db.Create(a).Error
}

// GET BY ID
func (r *answersRepository) GetByID(id uint) (*models.Answer, error) {
	var ans models.Answer
	err := r.db.First(&ans, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ans, err
}

// DELETE
func (r *answersRepository) Delete(id uint) error {
	return r.db.Delete(&models.Answer{}, id).Error
}
