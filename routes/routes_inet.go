package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v3 := api.Group("/v3")
	iF := v3.Group("/if")

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	dog.Get("/deleted", c.GetDeletedDogs)
	dog.Get("/half", c.GetDogHalf)

	//v1
	v1.Get("/:name", c.ParamsTest)
	v1.Get("fact/:num", c.Fact)
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/register", c.ValidTest)

	//v2
	v2.Get("/", c.HelloTestV2)

	//nickname
	iF.Get("/", c.TaxID)

}
