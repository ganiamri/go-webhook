package handlers

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//Execute executes docker-compose up to update system
func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pwd, _ := os.Getwd()
	dir := filepath.Join(pwd, ".", h.config.DirPath)

	parsed := strings.Split(r.URL.Path, "/")
	endpoint := parsed[len(parsed)-1]

	cmd := exec.Command(h.config.ProgramPath, h.config.EndPoint[endpoint].FilePath)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		log.Println("Error executing command: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("Execute command success at: ", time.Now().String())
	w.WriteHeader(http.StatusOK)
}
