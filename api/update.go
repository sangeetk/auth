package api

import (
	"time"
)

// UpdateRequest - Update the existing user
type UpdateRequest struct {
	AccessToken string    `json:"access_token"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Birthday    time.Time `json:"birthday"`
	Domain      string    `json:"domain"`
	Roles       []string  `json:"roles"`

	AddressType string `json:"address_type"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Zip         string `json:"zip"`

	Profession   string `json:"profession"`
	Introduction string `json:"introduction"`
}

// UpdateResponse - Returns error if update fails
type UpdateResponse struct {
	Err string `json:"err,omitempty"`
}
