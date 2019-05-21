package main

import (
	"github.com/ajanthan/istio-samples/services/echo/pkg/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/echo/{message}", handler.Echo)
	log.Fatal(http.ListenAndServe(":8080", router))
}
