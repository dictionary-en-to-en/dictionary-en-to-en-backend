package searchWord

import "github.com/gofiber/fiber/v2"

const RouteContext = "/search"

func RouteDecision(api fiber.Router) {
	apiGroup := api.Group(RouteContext)

	apiGroup.Post("", search)
}
