package models

import (
	"time"
)

type Twit struct {
	Id        uint64 ``
	UserId    uint64 `valid:"required,uint64"`
	Text      string `valid:"required,stringlength(5|255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
