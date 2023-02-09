package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"yoku.dev/repo/controllers"
	_ "yoku.dev/repo/migrations"
	"yoku.dev/repo/router"
)

func main() {
	fmt.Println("こんにちは, 世界!")

	visitController := controllers.VisitController{}
	cacheController := controllers.CacheController{}

	router.Post("/visit", visitController.Create)
	router.Get("/cache", cacheController.Index)
	router.Delete("/cache/{cache}", cacheController.Delete)

	for _, route := range router.Routes {
		fmt.Println(*route)
	}

	router.Listen(os.Getenv("APP_PORT"))
}
