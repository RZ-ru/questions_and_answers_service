package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"qa_service/internal/models"
	"qa_service/internal/repository"

	"github.com/go-chi/chi/v5"
)

type QuestionHandler struct {
	Repo *repository.Repository
}

func NewQuestionHandler(repo *repository.Repository) *QuestionHandler {
	return &QuestionHandler{Repo: repo}
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if input.Text == "" {
		http.Error(w, "text is required", http.StatusBadRequest)
		return
	}

	q := &models.Question{Text: input.Text}

	if err := h.Repo.Questions.Create(q); err != nil {
		http.Error(w, "failed to create", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(q)
}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.Repo.Questions.GetAll()
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(questions)
}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	q, err := h.Repo.Questions.GetByID(uint(id))
	if err != nil || q == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(q)
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Repo.Questions.Delete(uint(id)); err != nil {
		http.Error(w, "failed to delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
