package twit

import (
	"developit-golang-rest-api/api/v1/models"
	"developit-golang-rest-api/components/helpers/httperror"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/asaskevich/govalidator.v4"
	"net/http"
	"time"
)

type TwitController struct {
	db *gorm.DB
}

func NewTwitController(db *gorm.DB) *TwitController {
	return &TwitController{db}
}

func (controller TwitController) ActionCreate(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.Twit
	decoder.Decode(&input)

	twitModel := models.Twit{
		Text:      input.Text,
		UserId:    models.GetUserIdByToken(r, controller.db),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	_, validateError := govalidator.ValidateStruct(&twitModel)
	if validateError != nil {
		httperror.New(http.StatusUnprocessableEntity, validateError.Error(), rw)
		return
	}

	result := controller.db.Create(&twitModel)
	if len(result.GetErrors()) > 0 {
		httperror.New(http.StatusBadRequest, "can not create", rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

func (controller TwitController) ActionUpdate(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	decoder := json.NewDecoder(r.Body)
	var input models.Twit
	decoder.Decode(&input)

	twitModel := models.Twit{}

	controller.db.Where("id = ?", id).Find(&twitModel)

	twitModel.Text = input.Text

	_, validateError := govalidator.ValidateStruct(&twitModel)
	if validateError != nil {
		httperror.New(http.StatusUnprocessableEntity, validateError.Error(), rw)
		return
	}

	result := controller.db.Save(&twitModel)
	if len(result.GetErrors()) > 0 {
		httperror.New(http.StatusBadRequest, "can not create", rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

func (controller TwitController) ActionDelete(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	decoder := json.NewDecoder(r.Body)
	var input models.Twit
	decoder.Decode(&input)

	twitModel := models.Twit{}
	controller.db.Where("id = ?", id).Find(&twitModel)
	if twitModel.Id == 0 {
		httperror.New(http.StatusNotFound, "model not found", rw)
		return
	}
	result := controller.db.Delete(&twitModel)
	if len(result.GetErrors()) > 0 {
		httperror.New(http.StatusBadRequest, "can not create", rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
