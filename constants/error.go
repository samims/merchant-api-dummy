package constants

import "net/http"

// error constants
const (
	InternalServerError = "InternalServerError"
	BadRequest          = "BadRequest"
	Unauthorized        = "Unauthorized"
	NotFound            = "NotFound"
	Conflict            = "Conflict"
	UniqueEmailError    = "pq: duplicate key value violates unique constraint \"users_email_key\""
)

// ErrorString holds the string version of the error which is sent to the user
var ErrorString = map[string]string{
	InternalServerError: "Something went wrong",
	BadRequest:          "Bad request",
	Unauthorized:        "You are not authorized to perform this action",
	NotFound:            "Requested resource not found",
	Conflict:            "Resource already exists with this name",
	UniqueEmailError:    "Email should be unique",
}

// error code(response status code) constants
var ErrorCode = map[string]int{
	InternalServerError: http.StatusInternalServerError,
	BadRequest:          http.StatusBadRequest,
	Unauthorized:        http.StatusUnauthorized,
	NotFound:            http.StatusNotFound,
	Conflict:            http.StatusConflict,
	UniqueEmailError:    http.StatusBadRequest,
}
