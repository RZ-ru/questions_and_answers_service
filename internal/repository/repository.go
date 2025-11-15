package repository

import "gorm.io/gorm"

type Repository struct {
	Questions QuestionRepository
	Answers   AnswerRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Questions: NewQuestionRepository(db),
		Answers:   NewAnswerRepository(db),
	}
}
