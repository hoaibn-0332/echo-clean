package entity

import (
	"errors"
)

type Error struct {
	Message string `json:"error_message"`
	Code    int    `json:"error_code"`
}

type Errors struct {
	Errors     []Error
	HttpStatus int
}

var (
	E10001 = Error{Message: "Internal Server Error", Code: 10001}
	E10002 = Error{Message: "Not Found", Code: 10002}
	E10003 = Error{Message: "Conflict", Code: 10003}
	E10004 = Error{Message: "Bad Param Input", Code: 10004}
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
)
