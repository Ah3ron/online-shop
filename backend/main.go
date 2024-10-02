package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Создание нового шаблонизатора
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Пример маршрута для получения товаров
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

	// Маршрут для главной страницы
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Главная страница",
		})
	})

	// Маршрут для страницы с товарами
	app.Get("/products", func(c *fiber.Ctx) error {
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
		return c.Render("products", fiber.Map{
			"title":    "Список товаров",
			"products": products,
		})
	})

	// Запуск сервера
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
