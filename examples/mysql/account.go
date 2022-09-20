package main

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name string
}
