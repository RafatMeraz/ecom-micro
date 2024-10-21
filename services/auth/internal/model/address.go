package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId    uint
	Street    string
	Zip       string
	Area      string
	District  string
	Division  string
	IsDefault bool
}
