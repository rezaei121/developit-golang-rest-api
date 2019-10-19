package middlewares

import (
	"developit-golang-rest-api/api/modules/v1/user/models"
	"developit-golang-rest-api/components/helpers/httperror"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

type Authentication struct {
	db *gorm.DB
}

func NewAuthentication(db *gorm.DB) *Authentication {
	return &Authentication{db}
}

func (auth *Authentication) TokenAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "" {
			splitToken := strings.Split(token, "Bearer ")
			token = splitToken[1]
		}
		userTokenModel := models.UserToken{}
		auth.db.Where("token = ?", token).Find(&userTokenModel)
		if userTokenModel == (models.UserToken{}) {
			httperror.New(http.StatusUnauthorized, "Unauthorized! Your request was made with invalid credentials.", rw)
			return
		}
		next.ServeHTTP(rw, r)
	})
}
