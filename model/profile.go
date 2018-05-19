package model

import (
	"time"
)

type Profile struct {
	ID           uint64     `gorm:"primary_key"`
	Uid          uint64     `gorm:"uid"`
	Profession   string     `gorm:"profession"`
	Introduction string     `sql:"type:text"`
	CreatedAt    time.Time  `gorm:"created_at"`
	UpdatedAt    time.Time  `gorm:"updated_at"`
	DeletedAt    *time.Time `gorm:"deleted_at"`
}
