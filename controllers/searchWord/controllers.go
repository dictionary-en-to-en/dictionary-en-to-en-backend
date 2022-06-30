package searchWord

import (
	"DictionaryENtoENBackend/config/messages"
	inputforms "DictionaryENtoENBackend/controllers/_inputforms"
	"DictionaryENtoENBackend/tools"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func search(c *fiber.Ctx) error {

	// Parse inputForm.
	inputForm := new(inputforms.Search)
	if err := c.BodyParser(inputForm); err != nil {
		panic(errors.New(messages.InvalidInputForm))
	}

	// Validate inputForm.
	if err := inputForm.Validation(); err != nil {
		panic(err)
	}

	// send to api.
	output, err := tools.SendToDictionaryApi(*inputForm)
	if err != nil {
		tools.Sender(c, false, 200, nil, fiber.Map{
			"message": "Sorry pal, we couldn't find definitions for the word you were looking for",
		})
	}

	tools.Sender(c, true, 200, nil, output)
	return nil
}
