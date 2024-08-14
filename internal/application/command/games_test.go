package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"quiz/internal/application/dto"
	"quiz/internal/domain/entity"
	"quiz/internal/domain/errors"
	"quiz/internal/domain/repository"
	mocksRepository "quiz/internal/mocks/quiz/domain/repository"
)

func TestNewUserGameCommandHandler(t *testing.T) {
	tests := []struct {
		name               string
		questionRepository repository.QuestionRepository
		userGameRepository repository.UserGameRepository
		expectPanic        bool
	}{
		{
			name:               "panics when questionRepository is nil",
			questionRepository: nil,
			userGameRepository: new(mocksRepository.MockUserGameRepository),
			expectPanic:        true,
		},
		{
			name:               "panics when userGameRepository is nil",
			questionRepository: new(mocksRepository.MockQuestionRepository),
			userGameRepository: nil,
			expectPanic:        true,
		},
		{
			name:               "returns handler when both repositories are non-nil",
			questionRepository: new(mocksRepository.MockQuestionRepository),
			userGameRepository: new(mocksRepository.MockUserGameRepository),
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
				NewAddUserGameCommandHandler(tt.questionRepository, tt.userGameRepository)
			} else {
				handler := NewAddUserGameCommandHandler(tt.questionRepository, tt.userGameRepository)
				assert.IsType(t, &AddUserGameCommandHandler{}, handler)
			}
		})
	}
}

func TestUserGameCommandHandler_Handle(t *testing.T) {
	mockQuestionsRepo := new(mocksRepository.MockQuestionRepository)
	mockUserGameRepo := new(mocksRepository.MockUserGameRepository)

	h := NewAddUserGameCommandHandler(mockQuestionsRepo, mockUserGameRepo)

	command := AddUserGameCommand{
		Username: "testuser",
		Answers: []dto.UserAnswer{
			{
				QuestionID:       1,
				SelectedAnswerID: "A",
			},
			{
				QuestionID:       2,
				SelectedAnswerID: "B",
			},
		},
	}

	mockQuestionsRepo.On("GetByID", 1).Return(entity.Question{
		ID:              1,
		CorrectAnswerID: "A",
	}, nil)

	mockQuestionsRepo.On("GetByID", 2).Return(entity.Question{
		ID:              2,
		CorrectAnswerID: "B",
	}, nil)

	mockUserGameRepo.On("Save", mock.Anything).Return(nil)

	res, err := h.Handle(command)

	assert.NoError(t, err)
	assert.IsType(t, dto.UserGame{}, res)

	mockQuestionsRepo.AssertExpectations(t)
	mockUserGameRepo.AssertExpectations(t)
}

func TestUserGameCommandHandler_HandleWithValidationError(t *testing.T) {
	mockQuestionsRepo := new(mocksRepository.MockQuestionRepository)
	mockUserGameRepo := new(mocksRepository.MockUserGameRepository)

	h := NewAddUserGameCommandHandler(mockQuestionsRepo, mockUserGameRepo)

	command := AddUserGameCommand{
		Username: "testuser",
		Answers: []dto.UserAnswer{
			{
				QuestionID:       1,
				SelectedAnswerID: "F",
			},
		},
	}

	_, err := h.Handle(command)

	assert.IsType(t, errors.ValidationError{}, err)
}
