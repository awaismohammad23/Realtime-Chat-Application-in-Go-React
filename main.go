package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {

	app := fiber.New()

	// CORS
	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1743663",
		Key:     "b34b2a155f112acad972",
		Secret:  "a16574237a04461bd8c6",
		Cluster: "us2",
		Secure:  true,
	}

	app.Post("/api/messages", func(c fiber.Ctx) error {

		var data map[string]string

		err := pusherClient.Trigger("my-channel", "my-event", data)
		if err != nil {
			fmt.Println(err.Error())
		}

		return c.SendString("Hello, World! ")
	})

	//running backend on port 8000
	app.Listen(":8000")
}
