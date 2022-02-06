package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	m := make(map[string]int)

	app.Get("/getallscore", func(c *fiber.Ctx) error {
		return c.JSON(m)
	})

	app.Get("/getscorebyuserid/:id", func(c *fiber.Ctx) error {
		return c.JSON(m[c.Params("id")])
	})

	app.Post("/updatescoreplusone/:id", func(c *fiber.Ctx) error {
		userId := c.Params("id")
		if _, ok := m[userId]; ok {
			println(ok)
			m[userId] += 1
		} else {
			m[userId] = 1
		}
		return c.JSON(m)
	})

	app.Listen(":3000")

}
