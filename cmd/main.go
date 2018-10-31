package main

import (
	"log"
	"net/http"

	nconfig "github.com/ganiamri/go-webhook/config"
	nhandlers "github.com/ganiamri/go-webhook/handlers"

	"github.com/gorilla/mux"
)

const (
	// ConfigFileLocation is the file configuration of ths service.
	ConfigFileLocation = "webhook.yaml"
)

func main() {
	// Get Config
	configLoader := nconfig.NewYamlConfigLoader(ConfigFileLocation)
	config, err := configLoader.GetServiceConfig()
	if err != nil {
		log.Fatalf("Unable to load configuration: %v", err)
	}

	handler := nhandlers.NewHandler(config)

	r := mux.NewRouter()
	r.HandleFunc(config.EndPoint, handler.Execute).Methods("GET")

	if err := http.ListenAndServe(config.Address, r); err != nil {
		panic(err)
	}
}
