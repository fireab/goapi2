package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId uint   `json:"user_id"`
}
