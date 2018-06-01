package api

import (
	"time"
)

// RegisterRequest - New User registration request
type RegisterRequest struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
	Domain    string    `json:"domain"`
	Roles     []string  `json:"roles"`

	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`

	Profession   string `json:"profession"`
	Introduction string `json:"introduction"`
}

// RegisterResponse - New User registration response
type RegisterResponse struct {
	ConfirmToken string `json:"confirm_token,omitempty"`
	Err          string `json:"err,omitempty"`
}
