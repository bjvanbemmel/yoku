package models

import "gorm.io/gorm"

type AuthToken struct {
	gorm.Model
	UserID int
	User   User   `gorm:"type:varchar"`
	Key    string `gorm:"type:varchar"`
}
