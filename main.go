package main

import (
	"developit-golang-rest-api/api/v1/controllers/twit"
	"developit-golang-rest-api/api/v1/controllers/user"
	"developit-golang-rest-api/components/db"
	"developit-golang-rest-api/components/middlewares"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.RestApiHeader)

	userController := user.NewUserController(connection)
	apiRouter.HandleFunc("/v1/user/login", userController.ActionLogin).Methods("POST")
	apiRouter.HandleFunc("/v1/user/register", userController.ActionRegister).Methods("POST")

	// dashboard api
	authenticationMiddleware := middlewares.NewAuthentication(connection)
	dashboardRouter := router.PathPrefix("/api/v1/dashboard").Subrouter()
	dashboardRouter.Use(middlewares.RestApiHeader)
	dashboardRouter.Use(authenticationMiddleware.TokenAuthentication)
	dashboardRouter.HandleFunc("/user/profile", userController.ActionProfile).Methods("POST")

	twitController := twit.NewTwitController(connection)
	dashboardRouter.HandleFunc("/twit/create", twitController.ActionCreate).Methods("POST")
	dashboardRouter.HandleFunc("/twit/update/{id:[0-9]+}", twitController.ActionUpdate).Methods("PUT")
	dashboardRouter.HandleFunc("/twit/delete/{id:[0-9]+}", twitController.ActionDelete).Methods("DELETE")
	//swg api
	generateSwaggerAPI()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}

func generateSwaggerAPI() {
	var items []string
	var re = regexp.MustCompile(`(?ms){(.*)}`)
	err := filepath.Walk("./swagger-doc/api/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(info.Name()) == ".json" {
				fileData, _ := ioutil.ReadFile(path)
				res := re.ReplaceAllString(string(fileData), "${1}")
				items = append(items, res)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	server, _ := ioutil.ReadFile("./swagger-doc/server.json")
	result := strings.Replace(string(server), "\"paths\": {}", "\"paths\": {\n"+strings.Join(items, ","), -1) + "\n}"
	ioutil.WriteFile("./swagger-doc/doc.json", []byte(result), 0777)
}
