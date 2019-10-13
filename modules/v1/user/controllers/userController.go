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
	"gopkg.in/asaskevich/govalidator.v4"
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

func (controller UserController) ActionLogin(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var input models.UserLogin
	decoder.Decode(&input)
	user := &models.User{}
	controller.db.Where("email = ?", input.Email).Find(user)

	if user.Email == "" {
		httperror.New(http.StatusNotFound, "username or password is incorrect!", rw)
		return
	}

	if password.CheckHash(input.Password+user.Sult, user.Password) {
		userTokenModel := models.UserToken{
			UserId:    user.Id,
			Token:     randomstring.New(32),
			CreatedAt: time.Time{},
		}
		controller.db.Create(&userTokenModel)
		result, _ := json.Marshal(userTokenModel)
		rw.WriteHeader(http.StatusOK)
		rw.Write(result)
	} else {
		httperror.New(http.StatusBadRequest, "username or password is incorrect!", rw)
		return
	}
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

	_, validateError := govalidator.ValidateStruct(&userModel)
	if validateError != nil {
		httperror.New(http.StatusBadRequest, validateError.Error(), rw)
		return
	}

	result := controller.db.Create(&userModel)
	if len(result.GetErrors()) > 0 {
		httperror.New(http.StatusBadRequest, "can not create", rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
