package model

import "gorm.io/gorm"

// User has many CreditCards, UserID is the foreign key
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
