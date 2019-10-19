package main

import (
	"developit-golang-rest-api/api/modules/v1/user/controllers"
	"developit-golang-rest-api/components/db"
	"developit-golang-rest-api/components/middlewares"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	router := mux.NewRouter()
	router.Use(middlewares.RestApiHeader)

	userController := controllers.NewUserController(connection)
	router.HandleFunc("/api/v1/user/login", userController.ActionLogin).Methods("POST")
	router.HandleFunc("/api/v1/user/register", userController.ActionRegister).Methods("POST")

	// dashboard api
	authenticationMiddleware := middlewares.NewAuthentication(connection)
	dashboardRouter := router.PathPrefix("/api/v1/dashboard").Subrouter()
	dashboardRouter.Use(authenticationMiddleware.TokenAuthentication)
	dashboardRouter.HandleFunc("/user/profile", userController.ActionProfile).Methods("POST")

	//swg api
	router.HandleFunc("/api/swagger/api-doc", ActionApiDoc)
	router.HandleFunc("/api/swagger/schemas/{name}", ActionApiSchemas)

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}

func ActionApiDoc(rw http.ResponseWriter, r *http.Request) {
	file, error := ioutil.ReadFile("./api/swagger/api-doc.yaml")
	if error != nil {
		log.Fatal(error)
	}
	rw.Header().Set("Content-Type", "")
	rw.Header().Set("accept", "")
	fmt.Fprint(rw, string(file))
}

func ActionApiSchemas(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	file, error := ioutil.ReadFile("./api/swagger/schemas/" + params["name"] + ".yaml")
	if error != nil {
		log.Fatal(error)
	}
	rw.Header().Set("Content-Type", "")
	rw.Header().Set("accept", "")
	fmt.Fprint(rw, string(file))
}
