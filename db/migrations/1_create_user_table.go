package main

import (
	"developit-golang-rest-api/modules/v1/user/models"
	"github.com/jinzhu/gorm"
)

type CreateUserTableMigration struct {
	db *gorm.DB
}

func (m CreateUserTableMigration) Up() {
	m.db.CreateTable(&models.User{})
}

func (m CreateUserTableMigration) Down() {

}