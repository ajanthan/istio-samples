package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ajanthan/istio-samples/services/echo/pkg/model"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

//EchoFrontEnd calls EchoBackEnd
func EchoFrontEnd(url string) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		input := mux.Vars(req)["input"]
		reply, clientErr := http.Get(url + "/echo/" + input)
		if clientErr != nil {
			sendError(500, "Internal error", res)
			return
		}
		if reply.StatusCode == http.StatusOK {
			_, copyErr := io.Copy(res, reply.Body)
			if copyErr != nil {
				sendError(http.StatusInternalServerError, "Copy error", res)
				return
			}
		}
		sendError(reply.StatusCode, reply.Status, res)

	}
}

//EchoBackEnd echos the incoming message
func EchoBackEnd(version string) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		input := mux.Vars(req)["input"]
		message := model.Payload{}
		message.Message = input
		message.Version = version
		jsonEncoder := json.NewEncoder(res)
		err := jsonEncoder.Encode(message)
		if err != nil {
			fmt.Fprint(res, "Internal error")
			res.WriteHeader(500)
		}
	}
}

func sendError(status int, errorMessage string, res http.ResponseWriter) {
	fmt.Fprint(res, "Internal error")
	res.WriteHeader(status)
}
