package middlewares

import (
	"developit-golang-rest-api/helpers/httperror"
	"developit-golang-rest-api/modules/v1/user/models"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func TokenAuth(h httprouter.Handle, db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]

		userTokenModel := models.UserToken{}
		db.Where("token = ?", token).Find(&userTokenModel)
		if userTokenModel == (models.UserToken{}) {
			httperror.New(http.StatusUnauthorized, "Unauthorized! Your request was made with invalid credentials.", w)
			return
		}
		h(w, r, ps)
	}
}
