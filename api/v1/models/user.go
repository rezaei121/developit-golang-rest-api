package models

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Id        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(25);not null" valid:"required,stringlength(3|25)"`
	Surname   string `gorm:"type:varchar(35);not null" valid:"required,stringlength(3|35)"`
	Email     string `gorm:"type:varchar(255);unique;not null" valid:"email,required"`
	Password  string `gorm:"type:varchar(255);not null" valid:"required,stringlength(7|255)"`
	Status    int    `gorm:"type:int2;not null;default:1"`
	Sult      string `gorm:"type:varchar(8);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
	Email    string `valid:"email,required"`
	Password string `valid:"required,stringlength(7|255)"`
}

type UserToken struct {
	Id        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	UserId    uint64 `gorm:"unsigned;not null"`
	Token     string `gorm:"type:text;not null"`
	CreatedAt time.Time
}

func GetUserIdByToken(r *http.Request, db *gorm.DB) uint64 {
	token := r.Header.Get("Authorization")
	if token != "" {
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]
	}
	userTokenModel := UserToken{}
	db.Where("token = ?", token).Find(&userTokenModel)
	return userTokenModel.UserId
}
