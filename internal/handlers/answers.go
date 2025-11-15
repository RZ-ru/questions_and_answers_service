package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"qa_service/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type AnswerHandler struct {
	Service services.AnswerService
}

func NewAnswerHandler(s services.AnswerService) *AnswerHandler {
	return &AnswerHandler{Service: s}
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	qidStr := chi.URLParam(r, "id")
	qid, _ := strconv.Atoi(qidStr)

	var input struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Warn().
			Err(err).
			Msg("invalid json in CreateAnswer")
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	ans, err := h.Service.Create(uint(qid), input.UserID, input.Text)
	if err != nil {
		log.Warn().
			Err(err).
			Uint("question_id", uint(qid)).
			Str("user_id", input.UserID).
			Msg("failed to create answer")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Info().
		Uint("answer_id", ans.ID).
		Uint("question_id", ans.QuestionID).
		Msg("answer created")
	json.NewEncoder(w).Encode(ans)
}

func (h *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	ans, err := h.Service.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Uint("id", uint(id)).Msg("failed to get answer")
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if ans == nil {
		log.Warn().Uint("id", uint(id)).Msg("answer not found")
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ans)
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Service.Delete(uint(id)); err != nil {
		log.Error().
			Err(err).Uint("id", uint(id)).
			Msg("failed to delete answer")
		http.Error(w, "delete failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
