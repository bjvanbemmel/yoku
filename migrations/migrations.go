package migrations

import (
	. "yoku.dev/repo/database"
	"yoku.dev/repo/models"
)

func init() {
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.AuthToken{})
    Db.AutoMigrate(&models.VisitPath{})
	Db.AutoMigrate(&models.Visit{})
}
