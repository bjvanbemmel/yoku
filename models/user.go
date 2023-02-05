package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
    Name     string `gorm:"type:varchar"`
    Password string `gorm:"type:varchar"`
    Email    string `gorm:"type:varchar;index:idx_email,unique"`
}
