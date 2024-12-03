package controllers

import (
	"fmt"
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strings"
	"time"

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

	asciiString := strings.Join(asciiValues, " ")

	response := fmt.Sprintf("tax_id = %s : %s", tax, asciiString)

	return c.SendString(response)
}

func Register(c *fiber.Ctx) error {
	db := database.DBConn
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
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+/.[a-zA-Z]{2,}$`, fl.Field().String())
		return matched
	})

	validate.RegisterValidation("phone_number", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[0-9]{10}$`, fl.Field().String())
		return matched
	})

	validate.RegisterValidation("website", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^(https?://)?[a-z0-9.-]+\.[a-z]{2,}$`, fl.Field().String())
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

	db.Create(&user)
	return c.Status(201).JSON(user)
}

/////////////////////CRUD Dog//////////////////////////////////

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

func GetDogHalf(c *fiber.Ctx) error {
	db := database.DBConn
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id > ? && dog_id < ?", 50, 100)

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

func GetDogsSum(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	sum_red := 0
	sum_green := 0
	sum_pink := 0
	sum_nocolor := 0

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sum_red += 1
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sum_green += 1
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sum_pink += 1
		} else {
			typeStr = "no color"
			sum_nocolor += 1
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultColor{
		Data:        dataResults,
		Name:        "golang-test",
		Count:       len(dogs),
		Sum_Red:     sum_red,
		Sum_Green:   sum_green,
		Sum_Pink:    sum_pink,
		Sum_NoColor: sum_nocolor,
	}
	return c.Status(200).JSON(r)
}

// ///////////////////CRUD Company//////////////////////////////////

func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(company); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
	}

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func GetCompanys(c *fiber.Ctx) error {
	db := database.DBConn
	var companys []m.Company

	db.Find(&companys) //delelete = null
	return c.Status(200).JSON(companys)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// ///////////////////CRUD Profile//////////////////////////////////
func calculateAge(birthday string) (int, error) {
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, err
	}

	currentDate := time.Now()

	age := currentDate.Year() - birthDate.Year()
	if currentDate.YearDay() < birthDate.YearDay() {
		age--
	}

	return age, nil
}

func AddProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.Profile

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
	}

	var existingProfile m.Profile
	if err := db.Where("employee_id = ?", profile.EmployeeID).First(&existingProfile).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Employee ID already exists",
			"error":   fmt.Sprintf("Employee ID %d already exists", profile.EmployeeID),
		})
	}

	calculatedAge, err := calculateAge(profile.Birthday)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid birthday format",
			"error":   err.Error(),
		})
	}

	if profile.Age != 0 && profile.Age != calculatedAge {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Age does not match the provided birthday",
			"error":   fmt.Sprintf("Provided age: %d, calculated age: %d", profile.Age, calculatedAge),
		})
	}

	if profile.Age == 0 {
		profile.Age = calculatedAge
	}

	if err := db.Create(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not save profile to database",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(profile)
}

func GetProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	var profiles []m.Profile
	db.Find(&profiles)
	return c.Status(200).JSON(profiles)
}

func GetProfileSum(c *fiber.Ctx) error {
	db := database.DBConn
	var profile []m.Profile

	db.Find(&profile)
	gen_x := 0
	gen_y := 0
	gen_z := 0
	gen_b := 0
	gen_g := 0

	var dataResults []m.ProfileRes
	for _, v := range profile {
		GenStr := ""
		if v.Age < 24 {
			GenStr = "Gen Z"
			gen_z += 1
		} else if v.Age >= 24 && v.Age <= 41 {
			GenStr = "Gen Y"
			gen_y += 1
		} else if v.Age >= 42 && v.Age <= 56 {
			GenStr = "Gen X"
			gen_x += 1
		} else if v.Age >= 57 && v.Age <= 75 {
			GenStr = "Baby Boomer"
			gen_b += 1
		} else if v.Age > 75 {
			GenStr = "G.I. Generation"
			gen_g += 1
		} else {
			GenStr = "no gen"
		}

		d := m.ProfileRes{
			FirstName:  v.FirstName,
			EmployeeID: v.EmployeeID,
			Age:        v.Age,
			Gen:        GenStr,
		}

		dataResults = append(dataResults, d)
	}

	r := m.ResultGen{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(profile),
		GenX:  gen_x,
		GenY:  gen_y,
		GenZ:  gen_z,
		GenB:  gen_b,
		GenG:  gen_g,
	}
	return c.Status(200).JSON(r)
}

func UpdateProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.Profile
	id := c.Params("id")
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Where("id = ?", id).Updates(&profile)
	return c.Status(200).JSON(profile)
}

func RemoveProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var profile m.Profile
	result := db.Delete(&profile, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
