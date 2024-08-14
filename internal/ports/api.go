package ports

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"quiz/internal/application/command"
	"quiz/internal/application/dto"
	"quiz/internal/application/query"
	domainErrors "quiz/internal/domain/errors"
	"quiz/internal/ports/generated"
)

type Controller struct {
	questionProvider        *query.GetQuestionsQueryHandler
	userGamesCommandHandler *command.AddUserGameCommandHandler
	userGameQueryHandler    *query.GetUserGameQueryHandler
	userStatsQueryHandler   *query.GetUserStatsQueryHandler
}

func NewController(
	questionProvide *query.GetQuestionsQueryHandler,
	userGamesCommandHandler *command.AddUserGameCommandHandler,
	userGamesQueryHandler *query.GetUserGameQueryHandler,
	userStatsQueryHandler *query.GetUserStatsQueryHandler,
) Controller {
	return Controller{
		questionProvider:        questionProvide,
		userGamesCommandHandler: userGamesCommandHandler,
		userGameQueryHandler:    userGamesQueryHandler,
		userStatsQueryHandler:   userStatsQueryHandler,
	}
}

func (c Controller) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := c.questionProvider.Handle(query.GetQuestionsQuery{})

	questionsOutput := make([]generated.Question, len(questions))
	for i, question := range questions {

		options := make([]generated.Answer, len(question.Options))
		for i, answer := range question.Options {
			options[i] = generated.Answer{
				Id:   answer.ID,
				Text: answer.Text,
			}
		}

		questionsOutput[i] = generated.Question{
			CorrectAnswerId: question.CorrectAnswerID,
			Id:              question.ID,
			Options:         options,
			Text:            question.Text,
		}
	}

	jsonOutput, err := json.MarshalIndent(questionsOutput, "", "  ")
	if err != nil {
		http.Error(w, "Server internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonOutput)
}

func (c Controller) SubmitGame(w http.ResponseWriter, r *http.Request) {
	var request dto.UserGame
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	userGame, err := c.userGamesCommandHandler.Handle(
		command.AddUserGameCommand{
			Username: request.Username,
			Answers:  request.UserAnswer,
		})

	if errors.Is(err, domainErrors.ValidationError{}) {
		http.Error(w, fmt.Sprintf("Validation error. Message: %s", err.Error()), http.StatusBadRequest)
	}

	if errors.Is(err, domainErrors.QuestionNotFoundErr{}) {
		http.Error(w, fmt.Sprintf("Bad request. Message: %s", err.Error()), http.StatusBadRequest)
	}

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	jsonOutput, err := json.MarshalIndent(userGame, "", "  ")
	if err != nil {
		http.Error(w, "Server internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonOutput)
}

func (c Controller) GetGame(w http.ResponseWriter, r *http.Request, username string) {
	userGame, err := c.userGameQueryHandler.Handle(query.GetUserGameQuery{UserName: username})

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	gameOutput := generated.UserGame{
		Points:   userGame.Points,
		Username: username,
	}

	jsonOutput, err := json.MarshalIndent(gameOutput, "", "  ")
	if err != nil {
		http.Error(w, "Server internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonOutput)
}

func (c Controller) GetUserStats(w http.ResponseWriter, r *http.Request, username string) {
	stats, err := c.userStatsQueryHandler.Handle(query.GetUserStatsQuery{Username: username})

	if errors.Is(err, domainErrors.UserGameNotFoundErr{}) {
		http.Error(w, "User game not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Server internal error", http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		http.Error(w, "Server internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonOutput)
}
