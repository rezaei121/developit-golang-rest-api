package controllers

import (
	"developit-golang-rest-api/helpers/httperror"
	"developit-golang-rest-api/helpers/password"
	"developit-golang-rest-api/helpers/randomstring"
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

	sultString := randomstring.New(8)
	passwordHash := password.Hash(input.Password + sultString)

	userModel := models.User{
		Name:      input.Name,
		Surname:   input.Surname,
		Email:     input.Email,
		Password:  passwordHash,
		Status:    input.Status,
		Sult:      sultString,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	rw.Header().Set("Content-Type", "application/json")
	result := controller.db.Create(&userModel)
	if len(result.GetErrors()) > 0 {
		fmt.Println(result.GetErrors())
		httperror := httperror.New(http.StatusBadRequest, "can not create")
		errorMessage, _ := json.Marshal(&httperror)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(errorMessage)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
