package routes

import (
	c "go-fiber-test/controllers"
	mw "go-fiber-test/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InetRoutes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v3 := api.Group("/v3")
	iF := v3.Group("/if")

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", mw.BasicAuth(), c.GetDogs)
	dog.Get("/filter", mw.BasicAuth(), c.GetDog)
	dog.Get("/json", mw.BasicAuth(), c.GetDogsJson)
	dog.Post("/", mw.BasicAuth(), c.AddDog)
	dog.Put("/:id", mw.BasicAuth(), c.UpdateDog)
	dog.Delete("/:id", mw.BasicAuth(), c.RemoveDog)
	dog.Get("/deleted", mw.BasicAuth(), c.GetDeletedDogs)
	dog.Get("/half", mw.BasicAuth(), c.GetDogHalf)
	dog.Get("/sum", mw.BasicAuth(), c.GetDogsSum)

	//CRUD companys
	company := v1.Group("/company")
	company.Get("", mw.BasicAuth(), c.GetCompanys)
	company.Post("/", mw.BasicAuth(), c.AddCompany)
	company.Put("/:id", mw.BasicAuth(), c.UpdateCompany)
	company.Delete("/:id", mw.BasicAuth(), c.RemoveCompany)

	//CRUD profile
	profile := v1.Group("/profile")
	profile.Get("", c.GetProfiles)
	profile.Post("/", mw.BasicAuth(), c.AddProfile)
	profile.Put("/:id", mw.BasicAuth(), c.UpdateProfile)
	profile.Delete("/:id", mw.BasicAuth(), c.RemoveProfile)
	profile.Get("/result", c.GetProfileSum)

	//v1
	v1.Get("/:name", mw.BasicAuth(), c.ParamsTest)
	v1.Get("fact/:num", mw.BasicAuth(), c.Fact)
	v1.Get("/", mw.BasicAuth(), c.HelloTest)
	v1.Post("/", mw.BasicAuth(), c.BodyTest)
	v1.Post("/inet", mw.BasicAuth(), c.QueryTest)
	v1.Post("/register", mw.BasicAuth(), c.Register)

	//v2
	v2.Get("/", mw.BasicAuth(), c.HelloTestV2)

	//nickname
	iF.Get("/", mw.BasicAuth(), c.TaxID)

}
