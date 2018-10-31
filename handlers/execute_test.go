package handlers_test

import (
	"fmt"
	nconfig "ketitik/bitbucket-webhook/config"
	nhandlers "ketitik/bitbucket-webhook/handlers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
)

func TestExecuteOK(t *testing.T) {
	config := &nconfig.ServiceConfig{
		Address:     "localhost:8080",
		DirPath:     "../_scripts",
		ProgramPath: "/bin/sh",
		FilePath:    "hello.sh",
	}

	handler := nhandlers.NewHandler(config)

	r := mux.NewRouter()
	r.HandleFunc("/netmonk/execute", handler.Execute).Methods("GET")
	httpServer := httptest.NewServer(r)
	defer httpServer.Close()
	serverURL, _ := url.Parse(httpServer.URL)

	// Hit API Endpoint
	targetPath := fmt.Sprintf("%v/%v", serverURL, "/netmonk/execute")
	req, _ := http.NewRequest("GET", targetPath, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Unable to get worker status: %v", err)
	}
	defer resp.Body.Close()
}

func TestExecuteNOK(t *testing.T) {
	config := &nconfig.ServiceConfig{
		Address:     "localhost:8080",
		DirPath:     "_scripts",
		ProgramPath: "/bin/sh",
		FilePath:    "test_nok.sh",
	}

	handler := nhandlers.NewHandler(config)

	r := mux.NewRouter()
	r.HandleFunc("/netmonk/execute", handler.Execute).Methods("GET")
	httpServer := httptest.NewServer(r)
	defer httpServer.Close()
	serverURL, _ := url.Parse(httpServer.URL)

	// Hit API Endpoint
	targetPath := fmt.Sprintf("%v/%v", serverURL, "/netmonk/execute")
	req, _ := http.NewRequest("GET", targetPath, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Unable to get worker status: %v", err)
	}
	defer resp.Body.Close()
}
