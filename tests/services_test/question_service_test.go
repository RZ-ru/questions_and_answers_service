package services_test

import (
	"testing"

	"qa_service/internal/models"
	"qa_service/internal/services"

	"github.com/stretchr/testify/assert"
)

// ---- MOCK REPOSITORY ----

type mockQuestionRepo struct {
	CreateFunc func(q *models.Question) error
}

func (m *mockQuestionRepo) Create(q *models.Question) error {
	return m.CreateFunc(q)
}
func (m *mockQuestionRepo) GetAll() ([]models.Question, error) {
	return nil, nil
}
func (m *mockQuestionRepo) GetByID(id uint) (*models.Question, error) {
	return nil, nil
}
func (m *mockQuestionRepo) Delete(id uint) error {
	return nil
}

// ---- TESTS ----

func TestQuestionServiceCreate_EmptyText(t *testing.T) {
	repo := &mockQuestionRepo{}
	svc := services.NewQuestionService(repo)

	q, err := svc.Create("   ")

	assert.Nil(t, q)
	assert.EqualError(t, err, "text is required")
}

func TestQuestionServiceCreate_Success(t *testing.T) {
	repo := &mockQuestionRepo{
		CreateFunc: func(q *models.Question) error {
			q.ID = 10
			return nil
		},
	}
	svc := services.NewQuestionService(repo)

	q, err := svc.Create("Hello")

	assert.NoError(t, err)
	assert.Equal(t, uint(10), q.ID)
	assert.Equal(t, "Hello", q.Text)
}
