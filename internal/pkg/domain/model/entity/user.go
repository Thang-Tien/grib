package entity

import (
	"time"
)

type User struct {
	BaseEntity
	BirthDay    time.Time  `gorm:"column:birthday;type:datetime"`
	Email       string     `gorm:"column:email;type:varchar(255);not null"`
	PhoneNumber string     `gorm:"column:phone_number;type:varchar(255);not null"`
	Address     string     `gorm:"column:address;type:varchar(255);not null"`
	Customers   []Customer ``
	Wallet      Wallet     ``
}
