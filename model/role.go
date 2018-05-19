package model

import (
	"time"
)

type Role struct {
	ID        uint64     `gorm:"primary_key"`
	Uid       uint64     `gorm:"uid"`
	Roles     string     `sql:"type:text"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
