package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RunHTTPServer(createHandler func(app *fiber.App)) {
	apiRouter := fiber.New()
	setMiddlewares(apiRouter)

	rootRouter := fiber.New()

	rootRouter.Mount("/api", apiRouter)
	createHandler(apiRouter)

	err := rootRouter.Listen(fmt.Sprintf("%s:%s", viper.Get("HOST"), viper.Get("PORT")))
	if err != nil {
		log.Error(fmt.Sprintf("Error starting server, %s", err))
	}
}

func setMiddlewares(app *fiber.App) {
	addCorsMiddleware(app)
}

func addCorsMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: viper.Get("CORS_ALLOW_ORIGINS").(string),
	}))
}
