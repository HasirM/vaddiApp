package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/hasirm/vaddiapp/pkg/controllers"
)

func Setup(app *fiber.App) {
	//APIs
	app.Post("/api/validate", controllers.Validate)

	// JWT Middleware
	// check whether the ajwt is valid or not, if ajwt is expired and rjwt is still valid then
	// creates a new pair of JWT's.
	app.Use(controllers.Validate, jwtware.New(jwtware.Config{
		SigningKey:  []byte(controllers.SecretKey),
		TokenLookup: "cookie:ajwt",
	}))

}
