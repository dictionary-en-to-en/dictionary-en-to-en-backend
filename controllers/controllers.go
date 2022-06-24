package controllers

import (
	"DictionaryENtoENBackend/controllers/searchWord"
	"github.com/gofiber/fiber/v2"
)

func RegisterAll(api fiber.Router) {
	searchWord.RouteDecision(api)
}
