package error

import (
	"encoding/json"
)

// Error is object will return to client in case of error
type Error struct {
	Message string `json:"error_message"`
	Code    int    `json:"error_code"`
}

// ToJsonError convert error to json string
func (err *Error) ToJsonError(other ...Error) string {
	var errors []Error
	errors = append(errors, *err)
	if other != nil || len(other) > 0 {
		errors = append(errors, other...)
	}
	jsonData, e := json.Marshal(errors)
	if e != nil {
		return ""
	}

	return string(jsonData)
}

// Error implement error interface
var (
	// E10001 Internal Server Error
	E10001 = Error{Message: "Internal Server Error", Code: 10001}
	// E10002 Invalid related data.
	E10002 = Error{Message: "Invalid related data.", Code: 10002}
	// E10003 A required field cannot be empty.
	E10003 = Error{Message: "A required field cannot be empty.", Code: 10003}
	// E10004 The data does not meet validation requirements.
	E10004 = Error{Message: "The data does not meet validation requirements.", Code: 10004}
	// E10005 The requested data was not found.
	E10005 = Error{Message: "The requested data was not found.", Code: 10005}
	// E10009 The data already exists.
	E10009 = Error{Message: "The data is too long, please check again.", Code: 10009}
)

// Error implement error interface
var (
	// StatusBadRequestError Invalid input data.
	StatusBadRequestError = Error{Message: "Invalid input data.", Code: 40000}

	// UnsupportedMediaTypeError Unsupported media type.
	UnsupportedMediaTypeError = Error{Message: "Unsupported media type.", Code: 41500}
)
