package api

import (
	"time"
)

// UpdateRequest - Update the existing user
type UpdateRequest struct {
	Username  string    `json:"username"`
	FirstName string    `json:"fname"`
	LastName  string    `json:"lname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`

	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`

	Profession   string `json:"profession"`
	Introduction string `json:"introduction"`
}

// UpdateResponse - Returns error if update fails
type UpdateResponse struct {
	Err string `json:"err,omitempty"`
}
