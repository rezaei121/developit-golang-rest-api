package main

import (
	"developit-golang-rest-api/api/modules/v1/user/controllers"
	"developit-golang-rest-api/components/db"
	"developit-golang-rest-api/components/middlewares"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	router := httprouter.New()

	userController := controllers.NewUserController(connection)

	router.POST("/v1/user/login", userController.ActionLogin)
	router.POST("/v1/user/register", userController.ActionRegister)
	router.POST("/v1/dashboard/user/profile", middlewares.TokenAuth(userController.ActionProfile, connection))
	scanDir()
	router.GET("/api/swagger.yml", apiDoc)

	http.ListenAndServe(":8080", router)
}

func apiDoc(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, _ := ioutil.ReadFile("./swagger.yml")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(data))
}
