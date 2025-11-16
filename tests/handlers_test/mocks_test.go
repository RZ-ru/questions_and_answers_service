package handlers_test

import "qa_service/internal/models"

// ===== MOCK QuestionService =====

type mockQuestionService struct {
	GetByIDFunc func(id uint) (*models.Question, error)
}

func (m *mockQuestionService) Create(text string) (*models.Question, error) {
	return nil, nil
}

func (m *mockQuestionService) GetAll() ([]models.Question, error) {
	return nil, nil
}

func (m *mockQuestionService) GetByID(id uint) (*models.Question, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, nil
}

func (m *mockQuestionService) Delete(id uint) error {
	return nil
}

// ===== MOCK AnswerService =====

type mockAnswerService struct {
	GetByIDFunc func(id uint) (*models.Answer, error)
}

func (m *mockAnswerService) Create(questionID uint, userID, text string) (*models.Answer, error) {
	return nil, nil
}

func (m *mockAnswerService) GetByID(id uint) (*models.Answer, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, nil
}

func (m *mockAnswerService) Delete(id uint) error {
	return nil
}
