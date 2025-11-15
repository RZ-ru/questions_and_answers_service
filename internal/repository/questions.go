package repository

import (
	"errors"

	"qa_service/internal/models"

	"gorm.io/gorm"
)

type QuestionsRepository interface {
	Create(question *models.Question) error
	GetAll() ([]models.Question, error)
	GetByID(id uint) (*models.Question, error)
	Delete(id uint) error
}

type questionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) QuestionsRepository {
	return &questionsRepository{db: db}
}

// CREATE
func (r *questionsRepository) Create(q *models.Question) error {
	return r.db.Create(q).Error
}

// GET ALL
func (r *questionsRepository) GetAll() ([]models.Question, error) {
	var questions []models.Question
	err := r.db.Order("created_at DESC").Find(&questions).Error
	return questions, err
}

// GET BY ID + preload answers
func (r *questionsRepository) GetByID(id uint) (*models.Question, error) {
	var q models.Question
	err := r.db.
		Preload("Answers").
		First(&q, id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &q, err
}

// DELETE (cascade because of models)
func (r *questionsRepository) Delete(id uint) error {
	return r.db.Delete(&models.Question{}, id).Error
}
