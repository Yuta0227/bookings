package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// Custom configuration package
	"github.com/Yuta0227/bookings/internal/config"

	// Custom handlers package
	"github.com/Yuta0227/bookings/internal/handlers"

	// Custom template rendering package
	"github.com/Yuta0227/bookings/internal/render"

	// Session management package
	"github.com/alexedwards/scs/v2"
)

const portNumber = "localhost:8080" // The port number on which the server will run

// Declare global variables for application configuration and session management
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Set to true in production to enable production-specific configurations
	app.InProduction = false

	// Initialize session manager
	session = scs.New()

	// Set session lifetime to 24 hours
	session.Lifetime = 24 * time.Hour

	// Enable session cookie to persist even after the browser is closed
	session.Cookie.Persist = true

	// Set SameSite attribute to Lax to control cross-site cookie behavior
	session.Cookie.SameSite = http.SameSiteLaxMode

	// Ensure cookies are sent over HTTPS only in production
	session.Cookie.Secure = app.InProduction

	// Assign the session manager to the app configuration
	app.Session = session

	// Create a cache of parsed templates for rendering HTML
	templateCache, err := render.CreateTemplateCache()

	// Handle error if template cache creation fails
	if err != nil {
		fmt.Println("Error creating template cache", err)
		return // Exit the program
	}

	// Store the template cache in the app configuration
	app.TemplateCache = templateCache

	// Disable template caching (useful during development)
	app.UseCache = false

	// Initialize the repository with the app configuration
	repo := handlers.NewRepo(&app)

	// Set up handlers with the initialized repository
	handlers.NewHandlers(repo)

	// Pass the app configuration to the template rendering package
	render.NewTemplates(&app)

	// Create and configure the HTTP server
	srv := &http.Server{
		// Address and port the server will listen on
		Addr: portNumber,

		// Routes setup for the server
		Handler: routes(&app),
	}

	// Start the server and handle any errors
	err = srv.ListenAndServe()

	// Log the error and stop the program if the server fails to start
	log.Fatal(err)

	// Log a message indicating the server has started
	fmt.Printf("Starting application on port %s\n", portNumber)
}
