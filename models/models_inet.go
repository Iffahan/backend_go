package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Username     string `json:"username" validate:"required,min=3,max=32,username"`
	Email        string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Password     string `json:"password,omitempty" validate:"required,min=6,max=20"`
	LineID       string `json:"line_id,omitempty" validate:"required,min=3,max=32"`
	PhoneNumber  string `json:"phone_number,omitempty" validate:"required,min=10,max=13"`
	BusinessType string `json:"business_type,omitempty" validate:"required,min=3,max=32"`
	Website      string `json:"website,omitempty" validate:"required,min=3,max=32"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data  []DogsRes `json:"data"`
	Name  string    `json:"name"`
	Count int       `json:"count"`
}

type ResultColor struct {
	Data        []DogsRes `json:"data"`
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	Sum_Red     int       `json:"sum_red"`
	Sum_Green   int       `json:"sum_green"`
	Sum_Pink    int       `json:"sum_pink"`
	Sum_NoColor int       `json:"sum_nocolor"`
}

type Company struct {
	gorm.Model
	Name     string `json:"name"`
	Addresss string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	TaxID    string `json:"tax_id"`
	Website  string `json:"website"`
}

type Profile struct {
	gorm.Model
	EmployeeID int    `json:"employee_id,omitempty" validate:"required" gorm:"unique"`
	FirstName  string `json:"first_name,omitempty" validate:"required,min=3,max=32"`
	LastName   string `json:"last_name,omitempty" validate:"required,min=3,max=32"`
	Birthday   string `json:"birthday,omitempty" validate:"required" gorm:"type:date"`
	Age        int    `json:"age,omitempty" validate:"required,min=12,max=100"`
	Email      string `json:"email,omitempty" validate:"required,email"`
	Tel        string `json:"tel,omitempty" validate:"required,min=10,max=13"`
}

type ProfileRes struct {
	FirstName  string `json:"fullname"`
	EmployeeID int    `json:"employee_id"`
	Age        int    `json:"age"`
	Gen        string `json:"gen"`
}

type ResultProfile struct {
	Data  []ProfileRes `json:"data"`
	Name  string       `json:"name"`
	Count int          `json:"count"`
}

type ResultGen struct {
	Data  []ProfileRes `json:"data"`
	Name  string       `json:"name"`
	Count int          `json:"count"`
	GenX  int          `json:"gen_x"`
	GenY  int          `json:"gen_y"`
	GenZ  int          `json:"gen_z"`
	GenB  int          `json:"gen_b"`
	GenG  int          `json:"gen_g"`
}
