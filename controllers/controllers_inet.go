package controllers

import (
	"fmt"
	m "go-fiber-test/models"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Hello, World! V2")
}

func BodyTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.SendString("Hello " + name)
}

func Fact(c *fiber.Ctx) error {
	num, err := c.ParamsInt("num")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input. Please provide a non-negative integer.")
	}

	if num < 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Factorial is not defined for negative numbers.")
	}

	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	return c.SendString(fmt.Sprintf("%d! = %d", num, result))
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is " + a
	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}
