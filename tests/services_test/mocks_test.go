package services_test

import "qa_service/internal/models"

// ========== MOCK QUESTION REPOSITORY ==========

type mockQuestionRepo struct {
	CreateFunc  func(q *models.Question) error
	GetAllFunc  func() ([]models.Question, error)
	GetByIDFunc func(id uint) (*models.Question, error)
	DeleteFunc  func(id uint) error
}

func (m *mockQuestionRepo) Create(q *models.Question) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(q)
	}
	return nil
}

func (m *mockQuestionRepo) GetAll() ([]models.Question, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, nil
}

func (m *mockQuestionRepo) GetByID(id uint) (*models.Question, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, nil
}

func (m *mockQuestionRepo) Delete(id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

// ========== MOCK ANSWER REPOSITORY ==========

type mockAnswerRepo struct {
	CreateFunc  func(a *models.Answer) error
	GetByIDFunc func(id uint) (*models.Answer, error)
	DeleteFunc  func(id uint) error
}

func (m *mockAnswerRepo) Create(a *models.Answer) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(a)
	}
	return nil
}

func (m *mockAnswerRepo) GetByID(id uint) (*models.Answer, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, nil
}

func (m *mockAnswerRepo) Delete(id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}
