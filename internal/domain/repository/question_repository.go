package repository

import "quiz/internal/domain/entity"

type QuestionRepository interface {
	FindAll() []entity.Question
	GetByID(ID int) (entity.Question, error)
}
