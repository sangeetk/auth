package api

import (
	"time"
)

// UpdateRequest - Update the existing user
type UpdateRequest struct {
	AccessToken string `json:"access_token"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	// Email       string            `json:"email"`
	NewPassword string            `json:"new_password"`
	Birthday    time.Time         `json:"birthday"`
	Domain      string            `json:"domain"`
	Roles       []string          `json:"roles"`
	Address     Address           `json:"address"`
	Profile     map[string]string `json:"profile"`
}

// UpdateResponse - Returns error if update fails
type UpdateResponse struct {
	Err string `json:"err,omitempty"`
}
