package persistance

import (
	"quiz/internal/domain/entity"
	"quiz/internal/domain/errors"
)

type GameMemoryRepository struct {
	games map[string]*entity.UserGame
}

func NewGameMemoryRepository() *GameMemoryRepository {
	return &GameMemoryRepository{
		games: make(map[string]*entity.UserGame),
	}
}

func (repo *GameMemoryRepository) Save(userGame *entity.UserGame) {
	repo.games[userGame.Username] = userGame
}

func (repo *GameMemoryRepository) GetByUsername(username string) (*entity.UserGame, error) {
	game, ok := repo.games[username]
	if !ok {
		return nil, errors.UserGameNotFoundErr{Username: username}
	}

	return game, nil
}

func (repo *GameMemoryRepository) FindAll() map[string]*entity.UserGame {
	return repo.games
}
