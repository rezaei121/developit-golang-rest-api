package main

import (
	"developit-golang-rest-api/db"
	"developit-golang-rest-api/middlewares"
	"developit-golang-rest-api/modules/v1/user/controllers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	router := httprouter.New()

	userController := controllers.NewUserController(connection)

	router.POST("/v1/user/login", userController.ActionLogin)
	router.POST("/v1/user/register", userController.ActionRegister)
	router.POST("/v1/dashboard/user/profile", middlewares.TokenAuth(userController.ActionProfile, connection))

	http.ListenAndServe(":8080", router)
}
