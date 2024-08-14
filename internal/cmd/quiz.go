package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"quiz/internal/application/dto"
)

var apiUrl string

var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Run a simple quiz game",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		apiUrl = viper.GetString("API_URL")

		for {
			err := run()
			if err != nil {
				fmt.Print(err)
				//reader := bufio.NewReader(os.Stdin)
				//fmt.Print("\nPress Enter to restart the game or Ctrl+C to exit...")
				//_, _ = reader.ReadString('\n')
				//continue
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("\nPress Enter to restart the game or Ctrl+C to exit...")
			_, _ = reader.ReadString('\n')
		}
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)
}

func run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Printf("Your username is: %s \n", username)

	resp, err := http.Get(fmt.Sprintf("%s/game/%s", apiUrl, username))

	if err != nil {
		return fetchingGameError{}
	}

	defer resp.Body.Close()

	game := new(dto.UserGame)
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("No previous games found")
		game.Username = username
	} else {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return fetchingGameError{}
		}

		err = json.Unmarshal(body, game)
		if err != nil {
			return fetchingGameError{}
		}

		fmt.Print("\nYour previous result was: %d \n", game.Points)

	}

	fmt.Print("\nPress Enter to start a new game... \n")
	_, _ = reader.ReadString('\n')

	game.Points = 0
	game.UserAnswer = nil

	resp, err = http.Get(apiUrl + "/questions")

	if err != nil {
		return fetchingQuestionError{}
	}

	body, err := io.ReadAll(resp.Body)

	var questions []dto.Question
	err = json.Unmarshal(body, &questions)

	if err != nil {
		return fetchingQuestionError{}
	}

	for _, question := range questions {
		fmt.Printf("%s\n", question.Text)
		for _, option := range question.Options {
			fmt.Printf("%s) %s\n", option.ID, option.Text)
		}

		var choice string
		for {
			fmt.Print("Select an option: ")
			_, err := fmt.Scanf("%s\n", &choice)
			loweredChoice := strings.ToLower(choice)
			if err != nil || !checkOption(loweredChoice) {
				fmt.Println("Invalid choice, please try again. ")
				continue
			}

			if strings.ToLower(question.CorrectAnswerID) == loweredChoice {
				game.Points++

				fmt.Println("Correct answer.")
				fmt.Printf("Current score: %d \n", game.Points)
			} else {
				fmt.Printf("Incorrect. Correct answer is %s \n", question.CorrectAnswerID)
			}

			game.UserAnswer = append(game.UserAnswer, dto.UserAnswer{
				QuestionID:       question.ID,
				SelectedAnswerID: choice,
			})
			break
		}
	}

	err = postGame(game)
	if err != nil {
		return err
	}

	resp, err = http.Get(fmt.Sprintf("%s/user/%s/stats", apiUrl, username))

	if err != nil {
		return fetchingGameError{}
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	var userStats dto.UserStats
	err = json.Unmarshal(body, &userStats)
	if err != nil {
		return err
	}

	fmt.Printf("Your final score is: %d \n", game.Points)
	fmt.Printf("You were better than %v of all quizzers", userStats.RankScore)

	return nil
}

func checkOption(s string) bool {
	switch s {
	case "a", "b", "c", "d":
		return true
	default:
		return false
	}
}

func postGame(userGame *dto.UserGame) error {
	fmt.Println(userGame)
	jsonData, err := json.Marshal(userGame)

	if err != nil {
		return savingGameError{}
	}

	resp, err := http.Post(apiUrl+"/game", "application/json", bytes.NewBuffer(jsonData))
	if resp.StatusCode != http.StatusOK {
		return savingGameError{}
	}

	if resp.StatusCode != http.StatusOK {
		return savingGameError{}
	}

	return nil
}

type savingGameError struct {
}

func (savingGameError) Error() string {
	return fmt.Sprint("error while saving game data")
}

type fetchingQuestionError struct {
}

func (fetchingQuestionError) Error() string {
	return fmt.Sprint("error while fetching questions")
}

type fetchingGameError struct {
}

func (fetchingGameError) Error() string {
	return fmt.Sprint("error while fetching game")
}
