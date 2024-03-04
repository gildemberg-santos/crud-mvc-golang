package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname      string `gorm:"check:fullname <> ''" validate:"required"`
	TaxIdentifier string `gorm:"index:,unique;" validate:"required"`
	Email         string `gorm:"index:,unique;" validate:"required"`
	Password      string `gorm:"check:password <> ''" validate:"required" json:"-"`
}
