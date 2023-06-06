package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas-code42/url-shortner/api/routers"
	"github.com/lucas-code42/url-shortner/configs"
)

func InitializeServer() {
	app := fiber.New(fiber.Config{
		AppName:           "url-shortner",
		ServerHeader:      "api server",
		ColorScheme:       fiber.DefaultColors,
		EnablePrintRoutes: true,
	})
	
	routers.MountRouters(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", configs.SERVER_PORT)))
}
