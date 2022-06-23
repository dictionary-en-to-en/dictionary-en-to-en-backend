package pageNotFound

import "github.com/gofiber/fiber/v2"

func PageNotFound(c *fiber.Ctx) error {
	c.JSON(map[string]interface{}{
		"status":  false,
		"message": "Page Not Found",
	})
	return c.SendStatus(404)
}
