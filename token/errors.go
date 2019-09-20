package token

import (
	"errors"
)

// ErrorInvalidToken - Invalid Token
var ErrorInvalidToken = errors.New("Invalid Token")

// ErrorExpiredToken - Expired Token
var ErrorExpiredToken = errors.New("Expired Token")

// ErrorPasswordRecentlyUpdated - Recently recently updated
var ErrorPasswordRecentlyUpdated = errors.New("Password recently updated")
