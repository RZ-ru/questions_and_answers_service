package services

import (
	"errors"
	"strings"

	"qa_service/internal/models"
	"qa_service/internal/repository"

	"github.com/rs/zerolog/log"
)

type questionService struct {
	repo repository.QuestionRepository
}

func NewQuestionService(r repository.QuestionRepository) QuestionService {
	return &questionService{repo: r}
}

func (s *questionService) Create(text string) (*models.Question, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		log.Warn().
			Str("event", "create_question").
			Msg("empty question text")
		return nil, errors.New("text is required")
	}

	log.Info().
		Str("event", "create_question").
		Str("text", text).
		Msg("creating new question")

	q := &models.Question{Text: text}

	if err := s.repo.Create(q); err != nil {
		log.Error().
			Err(err).
			Str("event", "create_question").
			Msg("failed to save question in database")
		return nil, err
	}

	log.Info().
		Str("event", "create_question").
		Uint("created_id", q.ID).
		Msg("question created successfully")

	return q, nil
}

func (s *questionService) GetAll() ([]models.Question, error) {
	qs, err := s.repo.GetAll()
	if err != nil {
		log.Error().
			Err(err).
			Str("event", "get_all_questions").
			Msg("failed to load questions")
	}
	return qs, err
	//return s.repo.GetAll()
}

func (s *questionService) GetByID(id uint) (*models.Question, error) {
	q, err := s.repo.GetByID(id)
	if err != nil {
		log.Error().
			Err(err).
			Uint("id", id).
			Str("event", "get_question").
			Msg("database error")
		return nil, err
	}
	if q == nil {
		log.Warn().
			Uint("id", id).
			Str("event", "get_question").
			Msg("question not found")
	}
	return q, nil
}

func (s *questionService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		log.Error().
			Err(err).
			Uint("id", id).
			Str("event", "delete_question").
			Msg("failed to delete question")
	}
	return err
}
