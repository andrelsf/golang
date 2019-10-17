package models

import (
	"time"
)

type User struct {
	FullName	string		`json:"fullname"`
	Username	string		`json:"username"`
	Email		string		`json:"email"`
	Password	string		`json:"-"`
	Active		bool		`json:"active"`
	CreatedAt	time.Time	`json:"createat"`
	UpdatedAt	time.Time	`json:"updateat"`
}