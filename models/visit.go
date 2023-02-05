package models

import (
	"gorm.io/gorm"
)

type Visit struct {
    gorm.Model
    IP string `gorm:"type:varchar(45)"`
    URI string `gorm:"type:varchar"`
    UserAgent string `gorm:"type:varchar"`
}
