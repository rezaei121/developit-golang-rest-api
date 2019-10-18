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
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	router := httprouter.New()

	userController := controllers.NewUserController(connection)

	router.POST("/v1/user/login", userController.ActionLogin)
	router.POST("/v1/user/register", userController.ActionRegister)
	router.POST("/v1/dashboard/user/profile", middlewares.TokenAuth(userController.ActionProfile, connection))

	router.GET("/api/swagger/api-doc", ActionApiDoc)
	router.GET("/api/swagger/schemas/:name", ActionApiSchemas)

	//router.ServeFiles("/api/swagger/*filepath", http.Dir("api/swagger"))

	http.ListenAndServe(":8080", router)
}

func ActionApiDoc(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file, error := ioutil.ReadFile("./api/swagger/api-doc.yaml")
	if error != nil {
		log.Fatal(error)
	}
	rw.Header().Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(rw, string(file))
}

func ActionApiSchemas(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file, error := ioutil.ReadFile("./api/swagger/schemas/" + p.ByName("name") + ".yaml")
	if error != nil {
		log.Fatal(error)
	}
	rw.Header().Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(rw, string(file))
}
