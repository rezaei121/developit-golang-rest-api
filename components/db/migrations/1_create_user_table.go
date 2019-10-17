package main

import (
	"developit-golang-rest-api/api/modules/v1/user/models"
	"github.com/jinzhu/gorm"
)

type CreateUserTableMigration struct {
	db *gorm.DB
}

func (m CreateUserTableMigration) Up() {
	m.db.CreateTable(&models.User{})
	m.db.CreateTable(&models.UserToken{})
}

func (m CreateUserTableMigration) Down() {

}
