package jwt

import "fmt"

type ErrInvalidToken struct {
	token string
}

func (e ErrInvalidToken) Error() string {
	return fmt.Sprintf("invalid token %s", e.token)
}
