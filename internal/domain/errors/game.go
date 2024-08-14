package errors

import "fmt"

type UserGameNotFoundErr struct {
	Username string
}

func (e UserGameNotFoundErr) Error() string {
	return fmt.Sprintf("game not found for provided username: %s", e.Username)
}
