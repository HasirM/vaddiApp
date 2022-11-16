package main

import (
	// "fmt"
	// "encoding/json"
	// "io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "github.com/hasirm/vaddiapp/pkg/config"
	"github.com/hasirm/vaddiapp/pkg/database"
	"github.com/hasirm/vaddiapp/pkg/routes"
)

func main() {
	// fmt.Println("client server running....")

	// fmt.Println(config.Data().DBName)

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8001")

}
