package entity

type Question struct {
	ID              int
	Text            string
	Options         []Option
	CorrectAnswerID string
}

type Option struct {
	Text string
	ID   string
}
