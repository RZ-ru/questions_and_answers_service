package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"qa_service/internal/models"
	"qa_service/internal/repository"

	"github.com/go-chi/chi/v5"
)

type AnswerHandler struct {
	Repo *repository.Repository
}

func NewAnswerHandler(repo *repository.Repository) *AnswerHandler {
	return &AnswerHandler{Repo: repo}
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	qidStr := chi.URLParam(r, "id")
	qid, _ := strconv.Atoi(qidStr)

	var input struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if input.Text == "" || input.UserID == "" {
		http.Error(w, "fields required", http.StatusBadRequest)
		return
	}

	ans := &models.Answer{
		QuestionID: uint(qid),
		UserID:     input.UserID,
		Text:       input.Text,
	}

	if err := h.Repo.Answers.Create(ans); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(ans)
}

func (h *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	ans, err := h.Repo.Answers.GetByID(uint(id))
	if err != nil || ans == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ans)
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Repo.Answers.Delete(uint(id)); err != nil {
		http.Error(w, "delete failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
