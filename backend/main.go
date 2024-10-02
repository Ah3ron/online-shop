package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/", "./views")

	app.Get("/api/products", func(c *fiber.Ctx) error {
		products := []map[string]interface{}{
			{"id": 1, "name": "Product 1", "price": 100},
			{"id": 2, "name": "Product 2", "price": 200},
			{"id": 3, "name": "Product 3", "price": 300},
			{"id": 4, "name": "Product 4", "price": 400},
			{"id": 5, "name": "Product 5", "price": 500},
			{"id": 6, "name": "Product 6", "price": 600},
			{"id": 7, "name": "Product 7", "price": 700},
			{"id": 8, "name": "Product 8", "price": 800},
			{"id": 9, "name": "Product 9", "price": 900},
			{"id": 10, "name": "Product 10", "price": 1000},
		}
		return c.JSON(products)
	})

	app.Get("/:page?", renderPage)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func renderPage(c *fiber.Ctx) error {
	page := c.Path()[1:]
	if page == "" {
		page = "index"
	}
	return c.Render(page, fiber.Map{
		"title": page,
	})
}
