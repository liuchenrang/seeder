package error

import "fmt"

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ErrorMessage: %s", e.Code, e.Message)
}
func New(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}
