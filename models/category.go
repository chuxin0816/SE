package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(20);not null"`
}