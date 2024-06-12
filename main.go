package main

import (
	"fmt"

	"github.com/fuatanshori/gorm/config"
	"github.com/fuatanshori/gorm/database"
	"github.com/gofiber/fiber/v2"
)

func main(){

	database.DatabaseInit()
	app := fiber.New()
	app.Get("/",func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":"Hai",
		})
	})
	app.Listen(fmt.Sprintf(":%s",config.APP_PORT))
}