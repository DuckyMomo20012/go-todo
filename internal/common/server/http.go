package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func RunHttpServer(createHandler func(app *fiber.App)) {
	apiRouter := fiber.New()
	setMiddlewares(apiRouter)

	rootRouter := fiber.New()

	rootRouter.Mount("/api", apiRouter)
	createHandler(apiRouter)

	rootRouter.Listen(fmt.Sprintf("%s:%s", viper.Get("HOST"), viper.Get("PORT")))
}

func setMiddlewares(app *fiber.App) {
	addCorsMiddleware(app)
}

func addCorsMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: viper.Get("CORS_ALLOW_ORIGINS").(string),
	}))
}
