package main

import (
	"developit-golang-rest-api/api/modules/v1/user/controllers"
	"developit-golang-rest-api/components/db"
	"developit-golang-rest-api/components/middlewares"
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
	router.ServeFiles("/api/swagger/*filepath", http.Dir("api/swagger"))
	http.ListenAndServe(":8080", router)
}
