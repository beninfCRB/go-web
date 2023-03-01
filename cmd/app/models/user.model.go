package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type BindingUser struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
