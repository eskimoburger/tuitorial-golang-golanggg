package model

import "gorm.io/gorm"

type Todo struct {
	Task string
	gorm.Model
}
