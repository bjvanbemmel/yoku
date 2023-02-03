package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"yoku.dev/repo/router"
)

func main() {
	fmt.Println("こんにちは, 世界!")

	router.Get("/test", testCallback)
	router.Get("/api/{test}", testCallback)
	router.Get("/api/{hi}/test/{test}", testCallback)

    for _, route := range router.Routes {
        fmt.Println(*route)
    }

    router.Listen(os.Getenv("APP_PORT"))
}

func testCallback(c context.Context) {
    test := c.Value("test").(string)

	fmt.Println("Test callback!", test)
}
