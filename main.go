package main

import (
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello\n")
	})

	app.Post("/cli", func(c *fiber.Ctx) error {
		var command struct {
			Command string `json:"command"`
		}

		if err := c.BodyParser(&command); err != nil {
			return err
		}

		cmd := exec.Command("sh", "-c", command.Command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}

		return c.SendString(string(output))

	})

	app.Listen(":3000")
}
