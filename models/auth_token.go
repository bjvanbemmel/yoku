package models

import "gorm.io/gorm"

type AuthToken struct {
	gorm.Model
	UserID int
	User   User
	Key    string
}
