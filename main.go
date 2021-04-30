package main

import (
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func main() {
	app := fiber.New()

	app.Use(cache.New())

	// GET /john
	app.Get("/:email", func(c *fiber.Ctx) error {
		if !ValidateEmail(c.Params("email")) {
			return c.Status(204).Send(nil)
		}
		return c.Status(200).Send(nil)
	})

	log.Fatal(app.Listen(":8080"))
}

func ValidateEmail(email string) bool {
	log.Printf("Checking email: %s\n", email)
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !emailRegex.MatchString(email) {
		return false
	}
	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
