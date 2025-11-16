package services_test

import (
	"testing"

	"qa_service/internal/models"
	"qa_service/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestAnswerService_Create_EmptyText(t *testing.T) {
	svc := services.NewAnswerService(&mockAnswerRepo{}, &mockQuestionRepo{})

	a, err := svc.Create(1, "user1", "   ")

	assert.Nil(t, a)
	assert.EqualError(t, err, "text is required")
}

func TestAnswerService_Create_NoUserID(t *testing.T) {
	svc := services.NewAnswerService(&mockAnswerRepo{}, &mockQuestionRepo{})

	a, err := svc.Create(1, "", "hello")

	assert.Nil(t, a)
	assert.EqualError(t, err, "user_id is required")
}

func TestAnswerService_Create_QuestionDoesNotExist(t *testing.T) {
	svc := services.NewAnswerService(
		&mockAnswerRepo{},
		&mockQuestionRepo{
			GetByIDFunc: func(id uint) (*models.Question, error) {
				return nil, nil // вопрос не найден
			},
		},
	)

	a, err := svc.Create(5, "user1", "text")

	assert.Nil(t, a)
	assert.EqualError(t, err, "question does not exist")
}

func TestAnswerService_Create_Success(t *testing.T) {
	answerRepo := &mockAnswerRepo{
		CreateFunc: func(a *models.Answer) error {
			a.ID = 77
			return nil
		},
	}

	questionRepo := &mockQuestionRepo{
		GetByIDFunc: func(id uint) (*models.Question, error) {
			return &models.Question{ID: id}, nil
		},
	}

	svc := services.NewAnswerService(answerRepo, questionRepo)

	a, err := svc.Create(10, "user1", "hello")

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, uint(77), a.ID)
	assert.Equal(t, uint(10), a.QuestionID)
	assert.Equal(t, "user1", a.UserID)
}
