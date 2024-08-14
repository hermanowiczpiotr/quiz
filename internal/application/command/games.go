package command

import (
	"fmt"
	"strings"

	"quiz/internal/application/dto"
	"quiz/internal/domain/entity"
	"quiz/internal/domain/errors"
	"quiz/internal/domain/repository"
)

type AddUserGameCommand struct {
	Username string
	Answers  []dto.UserAnswer
}

type AddUserGameCommandHandler struct {
	questionsRepository repository.QuestionRepository
	userGameRepository  repository.UserGameRepository
}

func NewAddUserGameCommandHandler(
	questionRepository repository.QuestionRepository,
	userGameRepository repository.UserGameRepository,
) *AddUserGameCommandHandler {
	if questionRepository == nil {
		panic("questionRepository cannot be nil")
	}

	if userGameRepository == nil {
		panic("userGameRepository cannot be nil")
	}

	return &AddUserGameCommandHandler{
		questionsRepository: questionRepository,
		userGameRepository:  userGameRepository,
	}
}

func (handler AddUserGameCommandHandler) Handle(command AddUserGameCommand) (dto.UserGame, error) {
	game := &entity.UserGame{Username: command.Username}
	for _, answer := range command.Answers {
		answerID := strings.ToLower(answer.SelectedAnswerID)
		if !validateAnswer(answerID) {
			return dto.UserGame{}, errors.ValidationError{Message: fmt.Sprintf("Option should be a, b, c or d. %s provided", answerID)}
		}

		question, err := handler.questionsRepository.GetByID(answer.QuestionID)
		if err != nil {
			return dto.UserGame{}, err
		}

		userAnswer := entity.UserAnswer{
			QuestionID:       question.ID,
			SelectedAnswerID: answer.SelectedAnswerID,
		}

		if strings.ToLower(question.CorrectAnswerID) == answerID {
			game.AddPoint()
		}

		game.UserAnswer = append(game.UserAnswer, userAnswer)
	}

	handler.userGameRepository.Save(game)

	return dto.UserGame{
		Username:   game.Username,
		UserAnswer: command.Answers,
		Points:     game.Points,
	}, nil
}

func validateAnswer(answer string) bool {
	switch answer {
	case "a", "b", "c", "d":
		return true
	default:
		return false
	}
}
