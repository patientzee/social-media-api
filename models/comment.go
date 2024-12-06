package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text   string `json:"text"`
	UserID uint   `json:"user_id"`
	PostID uint   `json:"post_id"`
}
