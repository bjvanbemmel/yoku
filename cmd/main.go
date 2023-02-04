package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/joho/godotenv/autoload"
	. "yoku.dev/repo/database"
	_ "yoku.dev/repo/migrations"
	"yoku.dev/repo/models"
	"yoku.dev/repo/router"
)

func main() {
	fmt.Println("こんにちは, 世界!")

    router.Post("/auth", registerUser).Middleware(validateRegisterInput)
	router.Get("/test/{test}", testCallback).Middleware(testMiddleware)

	for _, route := range router.Routes {
		fmt.Println(*route)
	}

	router.Listen(os.Getenv("APP_PORT"))
}

func registerUser(c *router.Context) {
    pw, err := bcrypt.GenerateFromPassword(c.Value("password").([]byte), 14)
    if err != nil {
        c.WriteMap(map[string]any{
            "error": "Could not hash password.",
        }, 500)
    }

    user := models.User{
        Name: c.Value("name").(string),
        Email: c.Value("email").(string),
        Password: string(pw),
    }

    Db.Create(&user)

    Db.Create(&models.AuthToken{
        User: user,
        Key: uuid.NewString(),
    })
}

func validateRegisterInput(c *router.Context) error {
    req := c.Request.Body

    var b map[string]any
    json.NewDecoder(req).Decode(&b)

    if _, ok := b["name"]; ok == false {
        c.WriteMap(map[string]any{
            "error": "Please enter a name.",
        }, 422)

        return errors.New("Missing name.")
    }
    
    if _, ok := b["email"]; ok == false {
        c.WriteMap(map[string]any{
            "error": "Please enter an email.",
        }, 422)

        return errors.New("Missing email.")
    }
    
    if _, ok := b["password"]; ok == false {
        c.WriteMap(map[string]any{
            "error": "Please enter a password.",
        }, 422)

        return errors.New("Missing password.")
    }
    
    if _, ok := b["name"].(string); ok == false {
        c.WriteMap(map[string]any{
            "error": "Please enter a name.",
        }, 422)

        return errors.New("Missing name.")
    }
    
    if _, ok := b["email"].(string); ok == false {
        c.WriteMap(map[string]any{
            "error": "Please enter an email.",
        }, 422)

        return errors.New("Missing email.")
    }
    
    if _, ok := b["password"].(string); ok == false {
        c.WriteMap(map[string]any{
            "error": "Please enter a password.",
        }, 422)

        return errors.New("Missing password.")
    }

    c.WithValue("name", b["name"].(string))
    c.WithValue("email", b["email"].(string))
    c.WithValue("password", []byte(b["password"].(string)))

    return nil
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
