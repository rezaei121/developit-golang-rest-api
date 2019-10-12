package main

import (
	"developit-golang-rest-api/db"
	"developit-golang-rest-api/modules/v1/user/controllers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	router := httprouter.New()
	router.GET("/user/index", controllers.IndexAction)

	http.ListenAndServe(":8080", router)
}
