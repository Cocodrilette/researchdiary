package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	ID uint64 `gorm:"primaryKey;autoIncrement:true"`

	FirstName string
	LastName  string
}

func (a Author) FirstInitial() string {
	if len(a.FirstName) == 0 {
		return ""
	}
	return string([]rune(a.FirstName)[0])
}
