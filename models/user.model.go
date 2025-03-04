package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// 1 to many relation between user and posts
	Posts []Post `json:"posts" gorm:"foreignKey:UserId"`
}
