package controllers

import (
	"fmt"
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strings"

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

func TaxID(c *fiber.Ctx) error {
	tax := c.Query("tax_id")
	asciiValues := []string{}
	for _, v := range tax {
		asciiValues = append(asciiValues, fmt.Sprintf("%d", int(v)))
	}

	// Join the ASCII values into a single string
	asciiString := strings.Join(asciiValues, " ")

	// Construct the response string
	response := fmt.Sprintf("tax_id = %s : %s", tax, asciiString)

	return c.SendString(response) // Return plain text response
}

func ValidTest(c *fiber.Ctx) error {
	user := new(m.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
	}

	validate := validator.New()

	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, fl.Field().String())
		return matched
	})

	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, fl.Field().String())
		return matched
	})

	validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, fl.Field().String())
		return matched
	})

	validate.RegisterValidation("phone_number", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[0-9]{10}$`, fl.Field().String())
		return matched
	})

	validate.RegisterValidation("website", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^(https?://)?[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, fl.Field().String())
		return matched
	})

	if err := validate.Struct(user); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			validationErrors[field] = fmt.Sprintf("Field '%s' failed on the '%s' rule", field, tag)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validationErrors,
		})
	}

	return c.JSON(user)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDeletedDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var deletedDogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&deletedDogs)

	return c.Status(200).JSON(deletedDogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet
			DogID: v.DogID, //112
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}
