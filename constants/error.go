package constants

import "net/http"

// error constants
const (
	InternalServerError          = "InternalServerError"
	BadRequest                   = "BadRequest"
	Unauthorized                 = "Unauthorized"
	NotFound                     = "NotFound"
	Conflict                     = "Conflict"
	UniqueEmailError             = "pq: duplicate key value violates unique constraint \"users_email_key\""
	PaginationError              = "PaginationError"
	UserAlreadyPartOfAMerchant   = "UserAlreadyPartOfAMerchant"
	UserNotPartOfAnyMerchant     = "UserNotPartOfAnyMerchant"
	UserNotPartOfCurrentMerchant = "UserNotPartOfCurrentMerchant"
	UserIDIsRequired             = "UserIDIsRequired"
	ErrorUserIDOrEmailRequired   = "ErrorUserIDOrEmailRequired"
	UserNotFound                 = "UserNotFound"
	MerchantNotFound             = "MerchantNotFound"
	ErrorEmptyString             = "ErrorEmptyString"
	ErrorInvalidCredentials      = "ErrorInvalidCredentials"
	PermissionDenied             = "PermissionDenied"
)

// ErrorString holds the string version of the error which is sent to the user
var ErrorString = map[string]string{
	InternalServerError:          "Something went wrong",
	BadRequest:                   "Bad request",
	Unauthorized:                 "You are not authorized to perform this action",
	NotFound:                     "Requested resource not found",
	Conflict:                     "Resource already exists with this name",
	UniqueEmailError:             "Email should be unique",
	PaginationError:              "Pagination error kindly send correct pagination parameters",
	UserAlreadyPartOfAMerchant:   "User already part of a merchant",
	UserNotPartOfAnyMerchant:     "User not part of any merchant",
	UserNotPartOfCurrentMerchant: "User not part of current merchant",
	UserIDIsRequired:             "User id is required",
	ErrorUserIDOrEmailRequired:   "User id or email is required",
	UserNotFound:                 "User not found",
	MerchantNotFound:             "Merchant not found",
	ErrorEmptyString:             "The string cannot be empty",
	ErrorInvalidCredentials:      "Invalid credentials",
	PermissionDenied:             "Permission denied",
}

// error code(response status code) constants
var ErrorCode = map[string]int{
	InternalServerError:          http.StatusInternalServerError,
	BadRequest:                   http.StatusBadRequest,
	Unauthorized:                 http.StatusUnauthorized,
	NotFound:                     http.StatusNotFound,
	Conflict:                     http.StatusConflict,
	UniqueEmailError:             http.StatusBadRequest,
	PaginationError:              http.StatusBadRequest,
	UserAlreadyPartOfAMerchant:   http.StatusBadRequest,
	UserNotPartOfAnyMerchant:     http.StatusBadRequest,
	UserNotPartOfCurrentMerchant: http.StatusBadRequest,
	UserIDIsRequired:             http.StatusBadRequest,
	ErrorUserIDOrEmailRequired:   http.StatusInternalServerError,
	UserNotFound:                 http.StatusNotFound,
	MerchantNotFound:             http.StatusNotFound,
	ErrorEmptyString:             http.StatusBadRequest,
	ErrorInvalidCredentials:      http.StatusUnauthorized,
	PermissionDenied:             http.StatusForbidden,
}
