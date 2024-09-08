package main

import (
	"testing"

	"github.com/Yuta0227/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Errorf("type is not *chi.Mux, but is %T", v)
	}
}
