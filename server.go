package main

import (
	"DictionaryENtoENBackend/config"
	"DictionaryENtoENBackend/controllers"
	"DictionaryENtoENBackend/controllers/pageNotFound"
	middleware "DictionaryENtoENBackend/middleware/middlewarecors"
	"DictionaryENtoENBackend/middleware/slash"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	router := fiber.New()
	// use recover function for when server panic, recover server.
	router.Use(recover())
	router.Use(logger.New(logger.Config{TimeFormat: "2006-01-02 15:04:05"}))

	// Use middleware functions.
	router.Use(cors.New(middleware.Configs))
	router.Use(slash.Remover)

	api := router.Group("/api" + config.ApiVersion)
	controllers.RegisterAll(api)

	// when request sent by user and there's no handler, it will show 404 error.
	router.Use(pageNotFound.PageNotFound)

}
