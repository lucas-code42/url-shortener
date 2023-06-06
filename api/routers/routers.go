package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas-code42/url-shortner/api/routers/handlers"
)

func MountRouters(app *fiber.App) {
	app.Post("/url-shortner", handlers.UrlShortner)
	app.Get("/:url", handlers.UrlRedirect)
}
