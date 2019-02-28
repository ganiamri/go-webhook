package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	nconfig "github.com/ganiamri/go-webhook/config"
	nhandlers "github.com/ganiamri/go-webhook/handlers"

	"github.com/gorilla/mux"
)

func TestExecuteOK(t *testing.T) {
	config := &nconfig.ServiceConfig{
		Address:     "localhost:8080",
		DirPath:     "../_scripts",
		ProgramPath: "/bin/sh",
		EndPoint: map[string]nconfig.DirectoryFile{
			"test_001": nconfig.DirectoryFile{
				FilePath: "test-001.sh",
			},
		},
	}

	handler := nhandlers.NewHandler(config)

	r := mux.NewRouter()
	r.HandleFunc("/test_001", handler.Execute).Methods("GET")
	httpServer := httptest.NewServer(r)
	defer httpServer.Close()
	serverURL, _ := url.Parse(httpServer.URL)

	// Hit API Endpoint
	targetPath := fmt.Sprintf("%v/%v", serverURL, "/test_001")
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
		EndPoint: map[string]nconfig.DirectoryFile{
			"test_nok": nconfig.DirectoryFile{
				FilePath: "test_nok.sh",
			},
		},
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
