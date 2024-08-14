package repository

import "quiz/internal/domain/entity"

type UserGameRepository interface {
	Save(userGame *entity.UserGame)
	GetByUsername(username string) (*entity.UserGame, error)
	FindAll() map[string]*entity.UserGame
}
