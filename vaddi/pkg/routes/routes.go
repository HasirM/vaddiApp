package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/hasirm/vaddiapp/pkg/middleware"
)

func Setup(app *fiber.App) {
	//APIs
	// middleware := http.Post("http://127.0.0.1:8000/api/validate", controllers.Request)
	app.Post("api/validate", )

	// JWT Middleware
	// check whether the ajwt is valid or not, if ajwt is expired and rjwt is still valid then
	// creates a new pair of JWT's.
	app.Use(middleware.Validate, jwtware.New(jwtware.Config{
		SigningKey:  []byte(middleware.SecretKey),
		TokenLookup: "cookie:ajwt",
	}))

}
