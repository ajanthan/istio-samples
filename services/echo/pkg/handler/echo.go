package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Echo echos the incoming message
func Echo(res http.ResponseWriter, req *http.Request) {
	message := mux.Vars(req)["message"]
	fmt.Fprint(res, message)
}
