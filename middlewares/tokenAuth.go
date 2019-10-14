package middlewares

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func TokenAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("salam :)")
		h(w, r, ps)
	}
}
