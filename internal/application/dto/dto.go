package dto

type Question struct {
	ID              int      `json:"id"`
	Text            string   `json:"text"`
	Options         []Option `json:"options"`
	CorrectAnswerID string   `json:"correct_answer_id"`
}

type UserAnswer struct {
	QuestionID       int    `json:"question_id"`
	SelectedAnswerID string `json:"selected_answer_id"`
}

type Option struct {
	Text string `json:"text"`
	ID   string `json:"id"`
}

type UserGame struct {
	Username   string       `json:"username"`
	UserAnswer []UserAnswer `json:"user_answers"`
	Points     int          `json:"points"`
}

type UserStats struct {
	Username  string  `json:"username"`
	Points    int     `json:"points"`
	RankScore float64 `json:"rank_score"`
}
