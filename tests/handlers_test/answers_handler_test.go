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

func TestGetAnswerHandler(t *testing.T) {
	mockSvc := &mockAnswerService{
		GetByIDFunc: func(id uint) (*models.Answer, error) {
			return &models.Answer{
				ID:         id,
				QuestionID: 10,
				UserID:     "u123",
				Text:       "hello",
			}, nil
		},
	}

	h := handlers.NewAnswerHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/answers/3", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/answers/{id}", h.GetAnswer)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var ans models.Answer
	json.Unmarshal(w.Body.Bytes(), &ans)

	assert.Equal(t, uint(3), ans.ID)
	assert.Equal(t, uint(10), ans.QuestionID)
	assert.Equal(t, "u123", ans.UserID)
	assert.Equal(t, "hello", ans.Text)
}
