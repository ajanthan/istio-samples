package main

import (
	"github.com/ajanthan/istio-samples/services/echo/pkg/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	const defaultVersion = "1.0"
	const defaultProfile = "backend"
	const defaultBackendUrl = "http://echo-backend:8080"
	version, isVersionSet := os.LookupEnv("APP_VERSION")
	if !isVersionSet {
		log.Println("APP_VERSION is not set. Setting default version " + defaultVersion)
		version = defaultVersion
	}
	profile, isProfileSet := os.LookupEnv("APP_PROFILE")
	if !isProfileSet {
		log.Println("APP_PROFILE is not set. Setting default profile " + defaultProfile)
		profile = defaultProfile
	}
	backendUrl, isBackendUrlSet := os.LookupEnv("APP_BACKEND")
	if !isBackendUrlSet {
		log.Println("APP_BACKEND is not set. Setting default backend Url " + defaultBackendUrl)
		backendUrl = defaultBackendUrl
	}

	router := mux.NewRouter()
	if profile == defaultProfile {
		router.HandleFunc("/echo/{input}", handler.EchoBackEnd(version))
	} else {
		router.HandleFunc("/echo/{input}", handler.EchoFrontEnd(backendUrl))
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
