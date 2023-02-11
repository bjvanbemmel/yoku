package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Paginator struct {
    CurrentPage int
    PerPage int
}

func init() {
	var dsn string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Europe/London",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	Db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Paginate(page, perPage int) *gorm.DB {
    if page < 1 {
        page = 1
    }

    return Db.Offset((page - 1) * perPage).Limit(perPage)
}
