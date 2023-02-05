package models

import (
	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	IP        string `gorm:"type:varchar(45)"`
	VisitPathID     int
	VisitPath       VisitPath
	UserAgent string `gorm:"type:varchar"`
}
