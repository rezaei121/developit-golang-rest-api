package models

import "time"

type UserModel struct {
	Id int
	Name string
	Surname string
	Email string
	Password string
	Status int
	CreatedAt time.Time
	UpdatedAt time.Time
}
