package query

import (
	"quiz/internal/domain/entity"
	"quiz/internal/domain/repository"
)

type GetUserGameQuery struct {
	UserName string
}

type GetUserGameQueryHandler struct {
	userGameRepo repository.UserGameRepository
}

func NewGetUserGameQueryHandler(userGameRepo repository.UserGameRepository) *GetUserGameQueryHandler {
	if userGameRepo == nil {
		panic("userGameRepo cannot be nil")
	}

	return &GetUserGameQueryHandler{userGameRepo: userGameRepo}
}

func (handler GetUserGameQueryHandler) Handle(query GetUserGameQuery) (*entity.UserGame, error) {
	userGame, err := handler.userGameRepo.GetByUsername(query.UserName)
	if err != nil {
		return nil, err
	}

	return userGame, nil
}
