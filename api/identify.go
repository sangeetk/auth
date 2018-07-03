package api

import (
	"time"
)

// IdentifyRequest - Indentifies the user using JWT
type IdentifyRequest struct {
	AccessToken string `json:"access_token"`
}

// IdentifyResponse - Sends the user details as reponse
type IdentifyResponse struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	Err       string    `json:"err,omitempty"`
}
