package model

import (
	"gorm.io/gorm"
)

type Blog struct {
    gorm.Model
	Title         string `gorm:"column:title;unique;not null" json:"title"`
	Content	      string `gorm:"column:content;not null" json:"content"`
	CreatedUserID uint   `gorm:"column:created_user;not null" json:"created_user"`
}