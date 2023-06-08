package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas-code42/url-shortner/api/db"
)

func UrlRedirect(c *fiber.Ctx) error {
	alias := c.Params("alias")

	redis := db.MountRedis()
	defer redis.CloseDb()

	r, err := redis.Get(alias)
	if err != nil {
		return c.JSON(map[string]string{"msg": "cant find reference"})
	}

	return c.Redirect(r, 303)
}
