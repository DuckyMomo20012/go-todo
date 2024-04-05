package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func RunHttpServer(createHandler func(app *fiber.App)) {
	apiRouter := fiber.New()

	rootRouter := fiber.New()

	rootRouter.Mount("/api", apiRouter)
	createHandler(apiRouter)

	rootRouter.Listen(fmt.Sprintf("%s:%s", viper.Get("HOST"), viper.Get("PORT")))
}
