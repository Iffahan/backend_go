package routes

import (
	"go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")

	v1.Get("/", controllers.HelloTest)
	v2.Get("/", controllers.HelloTestV2)

	v1.Post("/", controllers.BodyTest)

	v1.Get("/:name", controllers.ParamsTest)

	v1.Post("/inet", controllers.QueryTest)

	v1.Post("/valid", controllers.ValidTest)

}
