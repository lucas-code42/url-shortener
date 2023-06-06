package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas-code42/url-shortner/api/db"
)

func UrlRedirect(c *fiber.Ctx) error {
	alias := c.Params("url")

	redis := db.MountRedis()
	defer redis.CloseDb()

	r, err := redis.Get(alias)
	if err != nil {
		return c.JSON(map[string]string{"msg": "error"})
	}

	return c.Redirect(r, 302)
}
