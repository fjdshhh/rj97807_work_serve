package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Uid      string
	Name     string
	Password string
	Role     int
	Email    string
}
