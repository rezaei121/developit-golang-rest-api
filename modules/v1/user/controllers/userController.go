package controllers

import (
	"developit-golang-rest-api/helpers/httperror"
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

	rw.Header().Set("Content-Type", "application/json")
	result := controller.db.Create(&userModel)
	if result.GetErrors() != nil {
		httperror := httperror.New(http.StatusBadRequest, "can not create")
		errorMessage, _ := json.Marshal(&httperror)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(errorMessage)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
