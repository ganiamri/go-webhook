package main

import (
	"fmt"
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
	for key := range config.EndPoint {
		r.HandleFunc(fmt.Sprintf("/%v", key), handler.Execute).Methods("GET")
	}

	log.Println("Running service on ", config.Address)
	if err := http.ListenAndServe(config.Address, r); err != nil {
		panic(err)
	}
}
