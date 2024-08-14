package errors

import "fmt"

type QuestionNotFoundErr struct {
	Id int
}

func (e QuestionNotFoundErr) Error() string {
	return fmt.Sprintf("question not found for id: %s", e.Id)
}
