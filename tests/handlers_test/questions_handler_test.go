package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"qa_service/internal/handlers"
	"qa_service/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

// ---- MOCK SERVICE ----

type mockQuestionService struct {
	GetByIDFunc func(id uint) (*models.Question, error)
}

func (m *mockQuestionService) Create(text string) (*models.Question, error) { return nil, nil }
func (m *mockQuestionService) GetAll() ([]models.Question, error)           { return nil, nil }
func (m *mockQuestionService) GetByID(id uint) (*models.Question, error) {
	return m.GetByIDFunc(id)
}
func (m *mockQuestionService) Delete(id uint) error { return nil }

// ---- TEST ----

func TestGetQuestionHandler(t *testing.T) {
	mockSvc := &mockQuestionService{
		GetByIDFunc: func(id uint) (*models.Question, error) {
			return &models.Question{ID: id, Text: "Test question"}, nil
		},
	}

	h := handlers.NewQuestionHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/questions/5", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/questions/{id}", h.GetQuestion)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp models.Question
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, uint(5), resp.ID)
	assert.Equal(t, "Test question", resp.Text)
}
