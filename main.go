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
	router.GET("/api/api-doc.json", apiDoc)

	http.ListenAndServe(":8080", router)
}

func scanDir() {
	var data string
	err := filepath.Walk("./api",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(info.Name()) == ".swg" {
				//fmt.Println(path, info.Size(), filepath.Ext(info.Name()))
				swgFile, _ := ioutil.ReadFile(path)
				data = data + string(swgFile) + "\n"
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile("./api/api-doc.json", []byte(data), 0644)
}

func apiDoc(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, _ := ioutil.ReadFile("./api/api-doc.json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(data))
}
