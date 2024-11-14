package error

import (
	"encoding/json"
)

type Error struct {
	Message string `json:"error_message"`
	Code    int    `json:"error_code"`
}

func (err *Error) ToJsonError(other ...Error) string {
	errs := append(other, *err)
	jsonData, e := json.Marshal(errs)
	if e != nil {
		return ""
	}

	return string(jsonData)
}

var (
	E10001 = Error{Message: "Internal Server Error", Code: 10001}
	E10002 = Error{Message: "Invalid related data.", Code: 10002}
	E10003 = Error{Message: "A required field cannot be empty.", Code: 10003}
	E10004 = Error{Message: "The data does not meet validation requirements.", Code: 10004}
	E10005 = Error{Message: "The requested data was not found.", Code: 10005}
	E10009 = Error{Message: "The data is too long, please check again.", Code: 10009}
)

var (
	// StatusBadRequestError Invalid input data.
	StatusBadRequestError = Error{Message: "Invalid input data.", Code: 40000}

	// UnsupportedMediaTypeError Unsupported media type.
	UnsupportedMediaTypeError = Error{Message: "Unsupported media type.", Code: 41500}
)
