package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"quiz/internal/application/command"
	"quiz/internal/application/query"
	domainService "quiz/internal/domain/service"
	"quiz/internal/infrastructure/persistance"
	"quiz/internal/ports"
	"quiz/internal/ports/generated"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run an api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		userGamesRepository := persistance.NewGameMemoryRepository()
		questionsRepository := persistance.NewQuestionMemoryRepository(persistance.PrepareQuestions())
		questionsProvider := query.NewGetQuestionsQueryHandler(questionsRepository)
		userGameCommandHandler := command.NewAddUserGameCommandHandler(questionsRepository, userGamesRepository)
		getUserGameQueryHandler := query.NewGetUserGameQueryHandler(userGamesRepository)
		userScoreComparisonService := domainService.NewUserScoreComparisonService(userGamesRepository)

		getUserStatsQueryHandler := query.NewGetUserStatsQueryHandler(userScoreComparisonService)
		controller := ports.NewController(
			questionsProvider,
			userGameCommandHandler,
			getUserGameQueryHandler,
			getUserStatsQueryHandler,
		)

		r := http.NewServeMux()

		h := generated.HandlerFromMux(controller, r)

		port := viper.GetString("PORT")
		s := &http.Server{
			Handler: h,
			Addr:    ":" + port,
		}

		go func() {
			if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("Could not listen on %s: %v\n", port, err)
			}

		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		log.Printf("App started on port: %s", port)
		<-quit

		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
