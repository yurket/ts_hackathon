package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	Name string
}
