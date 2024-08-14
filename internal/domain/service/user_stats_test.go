package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"quiz/internal/application/dto"
	"quiz/internal/domain/entity"
	errors2 "quiz/internal/domain/errors"
	"quiz/internal/domain/repository"
	"quiz/internal/mocks/quiz/domain/repository"
)

func TestNewUserScoreComparisonService(t *testing.T) {
	tests := []struct {
		name               string
		userGameRepository repository.UserGameRepository
		expectPanic        bool
	}{
		{
			name:               "panics when userGameRepo is nil",
			userGameRepository: nil,
			expectPanic:        true,
		},
		{
			name:               "good result",
			userGameRepository: new(mocks.MockUserGameRepository),
			expectPanic:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic but did not occur")
					}
				}()
				NewUserScoreComparisonService(tt.userGameRepository)
			} else {
				handler := NewUserScoreComparisonService(tt.userGameRepository)
				assert.IsType(t, &UserScoreComparisonService{}, handler)
			}
		})
	}
}

func TestCalculateScore(t *testing.T) {
	mockRepo := new(mocks.MockUserGameRepository)

	username := "testuser"
	userGame := entity.UserGame{
		Username: username,
		Points:   50,
	}

	allGames := map[string]*entity.UserGame{
		"user1":    {Username: "user1", Points: 40},
		"user2":    {Username: "user2", Points: 30},
		"user3":    {Username: "user3", Points: 60},
		"user4":    {Username: "user3", Points: 79},
		"user5":    {Username: "user3", Points: 29},
		"user6":    {Username: "user3", Points: 45},
		"testuser": {Username: "testuser", Points: 50},
	}

	mockRepo.On("GetByUsername", username).Return(&userGame, nil)
	mockRepo.On("FindAll").Return(allGames)

	service := NewUserScoreComparisonService(mockRepo)

	result, err := service.CalculateScore(username)

	assert.NoError(t, err)
	assert.Equal(t, username, result.Username)
	assert.Equal(t, 50, result.Points)
	assert.Equal(t, float64(67), result.RankScore)

	mockRepo.AssertExpectations(t)
}

func TestCalculateScore_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.MockUserGameRepository)

	username := "unknownuser"

	mockRepo.On("GetByUsername", username).Return(entity.UserGame{}, errors2.UserGameNotFoundErr{Username: username})

	service := NewUserScoreComparisonService(mockRepo)

	result, err := service.CalculateScore(username)

	assert.Error(t, err)
	assert.Equal(t, dto.UserStats{}, result)

	mockRepo.AssertExpectations(t)
}
