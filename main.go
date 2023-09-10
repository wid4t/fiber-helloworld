package main

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	ID       int    `json:"ID"`
	Messages string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		var r = Response{}
		r.Messages = "Hello, world"
		byteArray, _ := json.Marshal(r)
		return c.Send(byteArray)
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		var id = c.Query("id")
		var r = Response{}
		i, _ := strconv.Atoi(id)
		r.ID = i
		r.Messages = "Hello " + c.Params("name")
		byteArray, _ := json.Marshal(r)
		return c.Send(byteArray)
	})

	app.Listen(":3000")
}
