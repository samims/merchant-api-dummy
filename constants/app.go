package constants

const (
	JWT_EXPIRATION_DELTA_IN_MINUTES = 60 * 24
)

type ContextKey string

const (
	UserIDContextKey ContextKey = "UserID"
)
