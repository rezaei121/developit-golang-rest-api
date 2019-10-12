package controllers

import (
	"developit-golang-rest-api/modules/v1/user/models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db}
}

func IndexAction(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "user index action...")
}

func (controller UserController) ActionRegister(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var input models.User
	decoder.Decode(&input)

	userModel := models.User{
		Name:      input.Name,
		Surname:   input.Surname,
		Email:     input.Email,
		Password:  input.Password,
		Status:    input.Status,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	if controller.db.NewRecord(&userModel) {
		rw.WriteHeader(http.StatusNoContent)
	}

	rw.WriteHeader(http.StatusBadRequest)
}
