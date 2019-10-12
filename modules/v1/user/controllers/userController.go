package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewUserController() {

}

func IndexAction(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "user index action...")
}

func ActionRegister(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
