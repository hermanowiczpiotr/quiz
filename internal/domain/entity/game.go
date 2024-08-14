package entity

type UserGame struct {
	Username   string
	UserAnswer []UserAnswer
	Points     int
}

func (userGame *UserGame) AddPoint() {
	userGame.Points++
}

type UserAnswer struct {
	QuestionID       int
	SelectedAnswerID string
}
