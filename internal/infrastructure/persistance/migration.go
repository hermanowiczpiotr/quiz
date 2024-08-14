package persistance

import "quiz/internal/domain/entity"

func PrepareQuestions() []entity.Question {
	return []entity.Question{
		{
			ID:   1,
			Text: "What is the correct way to declare a variable in Go?",
			Options: []entity.Option{
				{
					Text: "var x int = 10",
					ID:   "A",
				},
				{
					Text: "x := 10",
					ID:   "B",
				},
				{
					Text: "int x = 10",
					ID:   "C",
				},
				{
					Text: "let x = 10",
					ID:   "D",
				},
			},
			CorrectAnswerID: "B",
		},
		{
			ID:   2,
			Text: "Which built-in function is used to determine the length of a slice in Go?",
			Options: []entity.Option{
				{
					Text: "size()",
					ID:   "A",
				},
				{
					Text: "count()",
					ID:   "B",
				},
				{
					Text: "len()",
					ID:   "C",
				},
				{
					Text: "length()",
					ID:   "D",
				},
			},
			CorrectAnswerID: "C",
		},
		{
			ID:   3,
			Text: "How do you initiate a new Goroutine in Go?",
			Options: []entity.Option{
				{
					Text: "start functionName()",
					ID:   "A",
				},
				{
					Text: "go functionName()",
					ID:   "B",
				},
				{
					Text: "run functionName()",
					ID:   "C",
				},
				{
					Text: "thread functionName()",
					ID:   "D",
				},
			},
			CorrectAnswerID: "B",
		},
		{
			ID:   4,
			Text: "What is the purpose of the defer statement in Go?",
			Options: []entity.Option{
				{
					Text: "To handle exceptions",
					ID:   "A",
				},
				{
					Text: "To schedule a function call to be run after the function completes",
					ID:   "B",
				},
				{
					Text: "To start a new Goroutine",
					ID:   "C",
				},
				{
					Text: "To import packages",
					ID:   "D",
				},
			},
			CorrectAnswerID: "B",
		},
	}
}
