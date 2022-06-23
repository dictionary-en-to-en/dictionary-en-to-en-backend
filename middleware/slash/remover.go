package slash

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"regexp"
	"strings"
)

// Remover is custom middleware to redirect routes with a trailing `/`.
func Remover(c *fiber.Ctx) error {
	originalUrl := utils.CopyString(c.OriginalURL())

	extMatch, err := regexp.MatchString("\\.[a-zA-Z0-9]+$", originalUrl)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(originalUrl, "/") && !extMatch {
		err := c.Redirect(originalUrl + "/")
		if err != nil {
			return err
		}

	}
	err = c.Next()
	if err != nil {
		return err
	}

	return nil
}
