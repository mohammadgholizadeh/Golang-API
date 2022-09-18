package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"primaryKey"`
	Pwd      string
	Id       int
	Role     string
	Email    string
	Phone    int
	Tickets  []Tickets `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	Message  []Message `gorm:"foreignkey:AdminId;association_foreignkey:Id"`
}

type Tickets struct {
	gorm.Model

	Id          int `gorm:"autoIncrement"`
	UserId      int
	Subject     string
	Message     string
	ReadByUser  int
	ReadByAdmin int
}

type Message struct {
	gorm.Model

	Id       int `gorm:"autoIncrement"`
	AdminId  int
	TicketId int
	Text     string
}
