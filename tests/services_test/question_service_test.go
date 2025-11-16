package services_test

import (
	"testing"

	"qa_service/internal/models"
	"qa_service/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestQuestionService_Create_EmptyText(t *testing.T) {
	repo := &mockQuestionRepo{}
	svc := services.NewQuestionService(repo)

	q, err := svc.Create("   ")

	assert.Nil(t, q)
	assert.EqualError(t, err, "text is required")
}

func TestQuestionService_Create_Success(t *testing.T) {
	repo := &mockQuestionRepo{
		CreateFunc: func(q *models.Question) error {
			q.ID = 42
			return nil
		},
	}

	svc := services.NewQuestionService(repo)

	q, err := svc.Create("hello")

	assert.NoError(t, err)
	assert.Equal(t, uint(42), q.ID)
	assert.Equal(t, "hello", q.Text)
}
