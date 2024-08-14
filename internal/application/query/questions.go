package query

import (
	"quiz/internal/domain/entity"
	"quiz/internal/domain/repository"
)

type GetQuestionsQuery struct {
}

type GetQuestionsQueryHandler struct {
	questionsRepository repository.QuestionRepository
}

func NewGetQuestionsQueryHandler(questionsRepository repository.QuestionRepository) *GetQuestionsQueryHandler {
	if questionsRepository == nil {
		panic("questionRepository cannot be nil")
	}

	return &GetQuestionsQueryHandler{
		questionsRepository: questionsRepository,
	}
}

func (provider GetQuestionsQueryHandler) Handle(_ GetQuestionsQuery) ([]entity.Question, error) {
	return provider.questionsRepository.FindAll(), nil
}
