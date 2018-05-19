package model

import (
	"time"
)

type Address struct {
	ID          uint64     `gorm:"primary_key"`
	Uid         uint64     `gorm:"Uid"`
	AddressType string     `gorm:"address_type"`
	Address1    string     `gorm:"address1"`
	Address2    string     `gorm:"address2"`
	City        string     `gorm:"city"`
	State       string     `gorm:"state"`
	Country     string     `gorm:"country"`
	Zip         string     `gorm:"zip"`
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   time.Time  `gorm:"updated_at"`
	DeletedAt   *time.Time `gorm:"deleted_at"`
}
