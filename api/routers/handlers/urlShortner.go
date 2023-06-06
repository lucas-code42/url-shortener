package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas-code42/url-shortner/api/db"
)

type dtoUrlShortner struct {
	Url      string `json:"url"`
	UrlAlias string `json:"urlAlias,omitempty"`
	NewUrl   string `json:"newUrl,omitempty"`
}

func UrlShortner(c *fiber.Ctx) error {
	url := new(dtoUrlShortner)
	if err := c.BodyParser(url); err != nil {
		return c.JSON(map[string]string{"msg": "err"})
	}

	redis := db.MountRedis()
	defer redis.CloseDb()

	err := redis.Create(url.Url, url.UrlAlias)
	if err != nil {
		return c.JSON(map[string]string{"msg": "err"})
	}

	data := dtoUrlShortner{
		Url:    url.Url,
		NewUrl: fmt.Sprintf("http://127.0.0.1:8080/%s", url.UrlAlias),
	}
	return c.JSON(data)
}
