package api

import (
	"time"
)

// ProfileRequest - Request the profile for the logged-in user
type ProfileRequest struct {
	AccessToken string `json:"access_token"`
}

// ProfileResponse - Returns the profile for the logged-in user
type ProfileResponse struct {
	Username      string              `json:"username"`
	Name          string              `json:"name"`
	FirstName     string              `json:"first_name"`
	LastName      string              `json:"last_name"`
	Email         string              `json:"email"`
	Birthday      time.Time           `json:"birthday"`
	InitialDomain string              `json:"initial_domain"`
	Roles         map[string][]string `json:"roles"`
	Address       Address             `json:"address"`
	Profile       Profile             `json:"profile"`
	Err           string              `json:"err,omitempty"`
}
