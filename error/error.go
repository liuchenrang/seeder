package error

import "fmt"

type Error struct {
	Code    int
	Message string
}
const ZK_NODE_EXITSTS = "zk: node already exists"
const SQL_NO_ROWS = "sql: no rows in result set"

func (e Error) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ErrorMessage: %s", e.Code, e.Message)
}
func New(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}
