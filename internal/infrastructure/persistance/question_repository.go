package persistance

import (
	"quiz/internal/domain/entity"
	"quiz/internal/domain/errors"
)

type QuestionMemoryRepository struct {
	questions []entity.Question
}

func NewQuestionMemoryRepository(questions []entity.Question) *QuestionMemoryRepository {
	return &QuestionMemoryRepository{questions: questions}
}

func (repo *QuestionMemoryRepository) FindAll() []entity.Question {
	return repo.questions
}

func (repo *QuestionMemoryRepository) GetByID(ID int) (entity.Question, error) {
	for _, question := range repo.questions {
		if question.ID == ID {
			return question, nil
		}
	}

	return entity.Question{}, errors.QuestionNotFoundErr{Id: ID}
}
