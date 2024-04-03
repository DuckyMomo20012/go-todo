package server

import "github.com/gofiber/fiber/v2"

func RunHttpServer(createHandler func(app *fiber.App)) {
	apiRouter := fiber.New()

	rootRouter := fiber.New()

	rootRouter.Mount("/api", apiRouter)
	createHandler(apiRouter)

	rootRouter.Listen(":8080")
}
