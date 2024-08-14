package query

import (
	"quiz/internal/application/dto"
	"quiz/internal/domain/service"
)

type GetUserStatsQuery struct {
	Username string
}

type GetUserStatsQueryHandler struct {
	userScoreProvider service.UserScoreComparisonServiceI
}

func NewGetUserStatsQueryHandler(userScoreProvider service.UserScoreComparisonServiceI) *GetUserStatsQueryHandler {
	if userScoreProvider == nil {
		panic("userScoreProvider cannot be nil")
	}

	return &GetUserStatsQueryHandler{userScoreProvider: userScoreProvider}
}

func (handler GetUserStatsQueryHandler) Handle(query GetUserStatsQuery) (dto.UserStats, error) {
	score, err := handler.userScoreProvider.CalculateScore(query.Username)
	if err != nil {
		return score, err
	}

	return score, nil
}
