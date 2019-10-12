package main

import (
	"developit-golang-rest-api/modules/v1/user/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main()  {
	router := httprouter.New()
	router.GET("/user/index", controllers.IndexAction)

	http.ListenAndServe(":8080", router)
}
