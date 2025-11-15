package repository

import "gorm.io/gorm"

type Repository struct {
	Questions QuestionsRepository
	Answers   AnswersRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Questions: NewQuestionsRepository(db),
		Answers:   NewAnswersRepository(db),
	}
}
