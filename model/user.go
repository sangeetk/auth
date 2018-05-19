package model

import (
	"time"
)

type User struct {
	ID       uint64    `gorm:"primary_key"`
	Username string    `gorm:"username,not null;unique"`
	Fname    string    `gorm:"fname"`
	Lname    string    `gorm:"lname"`
	Email    string    `gorm:"email,not null;unique"`
	Password []byte    `gorm:"password"`
	Birthday time.Time `gorm:"birthday"`

	// Source domain of registration
	Domain string `gorm:"domain"`

	// Confirm
	ConfirmToken string `gorm:"confirm_token"`
	Confirmed    bool   `gorm:"confirmed"`

	// Lock
	AttemptNumber int64     `gorm:"attempt_number"`
	AttemptTime   time.Time `gorm:"attempt_time"`
	Locked        time.Time `gorm:"locked"`

	// Recover
	RecoverToken       string    `gorm:"recover_token"`
	RecoverTokenExpiry time.Time `gorm:"recover_token_expiry"`

	//Timestamps
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
