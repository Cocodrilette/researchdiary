package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FirstName string
	LastName  string
}

func (a Author) FirstInitial() string {
	if len(a.FirstName) == 0 {
		return ""
	}
	return string([]rune(a.FirstName)[0])
}
