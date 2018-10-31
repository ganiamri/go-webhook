package handlers

import (
	nconfig "github.com/ganiamri/go-webhook/config"
)

// Handler holds the API endpoint's function handler.
type Handler struct {
	config *nconfig.ServiceConfig
}

// NewHandler function to make connection database into handler
func NewHandler(config *nconfig.ServiceConfig) *Handler {
	return &Handler{
		config: config,
	}
}
