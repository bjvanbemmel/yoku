package models

import "gorm.io/gorm"

type VisitPath struct {
    gorm.Model
    Path string `gorm:"type:varchar"`
}
