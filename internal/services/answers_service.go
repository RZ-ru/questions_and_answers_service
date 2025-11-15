package services

import (
	"errors"
	"strings"

	"qa_service/internal/models"
	"qa_service/internal/repository"

	"github.com/rs/zerolog/log"
)

type answerService struct {
	answers   repository.AnswerRepository
	questions repository.QuestionRepository
}

func NewAnswerService(a repository.AnswerRepository, q repository.QuestionRepository) AnswerService {
	return &answerService{
		answers:   a,
		questions: q,
	}
}

func (s *answerService) Create(questionID uint, userID, text string) (*models.Answer, error) {
	text = strings.TrimSpace(text)

	// Делаем валидацию
	if text == "" {
		log.Warn().
			Str("event", "create_answer").
			Uint("question_id", questionID).
			Msg("empty answer text")
		return nil, errors.New("text is required")
	}
	if userID == "" {
		log.Warn().
			Str("event", "create_answer").
			Uint("question_id", questionID).
			Msg("empty user_id")
		return nil, errors.New("user_id is required")
	}

	log.Info().
		Str("event", "create_answer").
		Uint("question_id", questionID).
		Str("user_id", userID).
		Msg("creating answer")

	// Проверка существования вопроса
	q, err := s.questions.GetByID(questionID)
	if err != nil {
		log.Error().
			Err(err).
			Str("event", "create_answer").
			Uint("question_id", questionID).
			Msg("failed to get question")
		return nil, errors.New("error receiving question")
	}
	if q == nil {
		log.Warn().
			Str("event", "create_answer").
			Uint("question_id", questionID).
			Msg("question does not exist")
		return nil, errors.New("question does not exist")
	}

	// Создаем объект ответа
	a := &models.Answer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
	}

	// Пытаемся сохранить ответ
	if err := s.answers.Create(a); err != nil {
		log.Error().
			Err(err).
			Str("event", "create_answer").
			Uint("question_id", questionID).
			Msg("failed to save answer to database")
		return nil, err
	}

	log.Info().
		Str("event", "create_answer").
		Uint("question_id", questionID).
		Uint("answer_id", a.ID).
		Msg("answer created successfully")
	return a, nil
}

func (s *answerService) GetByID(id uint) (*models.Answer, error) {
	ans, err := s.answers.GetByID(id)
	if err != nil {
		log.Error().
			Err(err).
			Uint("id", id).
			Str("event", "get_answer").
			Msg("database error")
		return nil, err
	}
	if ans == nil {
		log.Warn().
			Uint("id", id).
			Str("event", "get_answer").
			Msg("answer not found")
	}
	return ans, nil
}

func (s *answerService) Delete(id uint) error {
	err := s.answers.Delete(id)
	if err != nil {
		log.Error().
			Err(err).
			Uint("id", id).
			Str("event", "delete_answer").
			Msg("failed to delete answer")
	}
	return err
}
