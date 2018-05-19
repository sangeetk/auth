package api

import (
	"time"
)

type RegisterRequest struct {
	Domain   string    `json:"domain"`
	Username string    `json:"username"`
	Fname    string    `json:"fname"`
	Lname    string    `json:"lname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`

	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`

	Profession   string `json:"profession"`
	Introduction string `json:"introduction"`
}

type RegisterResponse struct {
	ConfirmToken string `json:"confirm_token,omitempty"`
	Err          string `json:"err,omitempty"`
}
