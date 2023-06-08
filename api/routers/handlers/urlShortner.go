package handlers

import (
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lucas-code42/url-shortner/api/db"
)

type dtoUrlShortner struct {
	Url    string `json:"url"`
	NewUrl string `json:"newUrl,omitempty"`
}

func (dto *dtoUrlShortner) validateUrl() bool {
	regex := regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
	return regex.MatchString(dto.Url)

}

func UrlShortner(c *fiber.Ctx) error {
	url := new(dtoUrlShortner)
	if err := c.BodyParser(url); err != nil {
		c.Status(500)
		return c.JSON(map[string]string{"msg": "could not parse serialize object"})
	}

	if !url.validateUrl() {
		c.Status(400)
		return c.JSON(map[string]string{"msg": "invalid url"})
	}

	redis := db.MountRedis()
	defer redis.CloseDb()

	alias := uuid.NewString()
	err := redis.Create(url.Url, alias)
	if err != nil {
		c.Status(500)
		return c.JSON(map[string]string{"msg": "cannot create reference"})
	}

	data := dtoUrlShortner{
		Url:    url.Url,
		NewUrl: fmt.Sprintf("http://127.0.0.1:8080/%s", alias),
	}
	c.Status(200)
	return c.JSON(data)
}
