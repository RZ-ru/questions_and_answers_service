package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"qa_service/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type QuestionHandler struct {
	Service services.QuestionService
}

func NewQuestionHandler(s services.QuestionService) *QuestionHandler {
	return &QuestionHandler{Service: s}
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Warn().
			Err(err).
			Msg("invalid json in CreateQuestion")
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	q, err := h.Service.Create(input.Text)
	if err != nil {
		log.Warn().
			Err(err).
			Str("text", input.Text).
			Msg("failed to create question")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Info().
		Uint("id", q.ID).
		Msg("question created")
	json.NewEncoder(w).Encode(q)
}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.Service.GetAll()
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to load questions")
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(questions)
}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	q, err := h.Service.GetByID(uint(id))
	if err != nil {
		log.Error().
			Err(err).Uint("id", uint(id)).
			Msg("failed to get question")
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if q == nil {
		log.Warn().
			Uint("id", uint(id)).
			Msg("question not found")
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(q)
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Service.Delete(uint(id)); err != nil {
		log.Error().Err(err).
			Uint("id", uint(id)).
			Msg("failed to delete question")
		http.Error(w, "failed to delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
