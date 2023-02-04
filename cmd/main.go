package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "yoku.dev/repo/migrations"
	"yoku.dev/repo/router"
)

func main() {
	fmt.Println("こんにちは, 世界!")

	router.Get("/test/{test}", testCallback).Middleware(testMiddleware)

	for _, route := range router.Routes {
		fmt.Println(*route)
	}

	router.Listen(os.Getenv("APP_PORT"))
}

func testCallback(c *router.Context) {
	test := c.Value("test").(string)

	c.WriteMap(map[string]any{
		"status": test,
	}, 200)
}

func testMiddleware(c *router.Context) error {
	test := c.Context.Value("test").(string)

	fmt.Println("Test middleware!", test)

	return nil
}
