package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	. "yoku.dev/repo/database"
	_ "yoku.dev/repo/migrations"
	"yoku.dev/repo/models"
	"yoku.dev/repo/router"
)

func main() {
	fmt.Println("こんにちは, 世界!")

	router.Post("/visit", registerVisit)

	for _, route := range router.Routes {
		fmt.Println(*route)
	}

	router.Listen(os.Getenv("APP_PORT"))
}

type VisitBody struct {
	URL string
}

func registerVisit(c *router.Context) {
	agent := c.Request.UserAgent()

	var vb VisitBody
	json.NewDecoder(c.Request.Body).Decode(&vb)

	var ip string
	ip = c.Request.Header.Get("X-Real-IP")

	if ip == "" {
		var err error
		ip, _, err = net.SplitHostPort(c.Request.RemoteAddr)

		if err != nil {
			c.WriteMap(map[string]any{
				"error": "Could not parse IP address.",
			}, 500)

			return
		}
	}

	if vb.URL == "" {
		c.WriteMap(map[string]any{
			"error":   "missing_field",
			"message": "The `url` field must be filled.",
		}, 422)

		return
	}

	Db.Create(&models.Visit{
		UserAgent: agent,
		URI:       vb.URL,
		IP:        ip,
	})
}
