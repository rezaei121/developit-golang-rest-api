package models

import "time"

type User struct {
	Id        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(25);not null"`
	Surname   string `gorm:"type:varchar(35);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Status    int    `gorm:"type:int2;not null;default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
