package service

import (
	"math"

	"quiz/internal/application/dto"
	"quiz/internal/domain/repository"
)

type UserScoreComparisonServiceI interface {
	CalculateScore(username string) (dto.UserStats, error)
}

type UserScoreComparisonService struct {
	gamesRepo repository.UserGameRepository
}

func NewUserScoreComparisonService(userGameRepository repository.UserGameRepository) *UserScoreComparisonService {
	if userGameRepository == nil {
		panic("userGameRepository cannot be nil")
	}

	return &UserScoreComparisonService{gamesRepo: userGameRepository}
}

func (provider UserScoreComparisonService) CalculateScore(username string) (dto.UserStats, error) {
	userGame, err := provider.gamesRepo.GetByUsername(username)
	if err != nil {
		return dto.UserStats{}, err
	}

	games := provider.gamesRepo.FindAll()
	totalGames := len(games)

	worstResultsCounter := 0
	for _, game := range games {
		if game.Username == userGame.Username {
			continue
		}

		if game.Points < userGame.Points {
			worstResultsCounter++
		}
	}

	var rankScore float64
	if totalGames > 1 {
		rankScore = math.Round(float64(worstResultsCounter) / float64(totalGames-1) * 100)
	} else {
		rankScore = 100.0
	}

	return dto.UserStats{
		Username:  username,
		Points:    userGame.Points,
		RankScore: rankScore,
	}, nil
}
