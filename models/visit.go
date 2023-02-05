package models

import (
	"net"

	"gorm.io/gorm"
)

type Visit struct {
    gorm.Model
    IP net.IP
    URI string `gorm:"type:varchar"`
    UserAgent string `gorm:"type:varchar"`
}
