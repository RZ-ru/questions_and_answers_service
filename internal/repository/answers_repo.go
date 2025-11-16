package repository

import (
	"errors"

	"qa_service/internal/models"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(answer *models.Answer) error
	GetByID(id uint) (*models.Answer, error)
	Delete(id uint) error
}

type answersRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answersRepository{db: db}
}

func (r *answersRepository) Create(a *models.Answer) error {
	return r.db.Create(a).Error
}

func (r *answersRepository) GetByID(id uint) (*models.Answer, error) {
	var ans models.Answer
	err := r.db.First(&ans, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ans, err
}

func (r *answersRepository) Delete(id uint) error {
	return r.db.Delete(&models.Answer{}, id).Error
}
