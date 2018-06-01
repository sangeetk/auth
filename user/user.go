package user

import (
	"time"
)

// Profile - profile fields
type Profile struct {
	Profession   string `json:"profession"`
	Introduction string `json:"introduction"`
}

// Address - address fields
type Address struct {
	AddressType string `json:"address_type"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Zip         string `json:"zip"`
}

// User - user fields
type User struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  []byte    `json:"password"`
	Birthday  time.Time `json:"birthday"`

	// Source domain of registration
	InitialDomain string `json:"initial_domain"`

	// Roles for each domain
	Roles map[string][]string `json:"roles"`

	// Profile
	Profile Profile `json:"profile"`

	// Address
	Address Address `json:"address"`

	// Confirm
	ConfirmToken string `json:"confirm_token"`
	Confirmed    bool   `json:"confirmed"`

	// Lock
	AttemptNumber int64     `json:"attempt_number"`
	AttemptTime   time.Time `json:"attempt_time"`
	Locked        time.Time `json:"locked"`

	// Recover
	RecoverToken       string    `json:"recover_token"`
	RecoverTokenExpiry time.Time `json:"recover_token_expiry"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
