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
