package main

import (
	"net/http"

	// Importing custom configuration package
	"github.com/Yuta0227/bookings/pkg/config"
	// Importing custom handlers package
	"github.com/Yuta0227/bookings/pkg/handlers"
	// Importing chi router for routing HTTP requests
	"github.com/go-chi/chi"
	// Importing middleware from chi for additional functionality
	"github.com/go-chi/chi/middleware"
)

// routes sets up the HTTP routes for the application
// It takes a pointer to AppConfig as an argument to access the application configuration
func routes(app *config.AppConfig) http.Handler {
	
	// Initialize a new chi router (multiplexer)
	mux := chi.NewRouter()

	// Recover from panics in HTTP handlers and respond with a 500 error
	mux.Use(middleware.Recoverer)
	// Custom middleware to prevent Cross-Site Request Forgery (CSRF) attacks
	mux.Use(NoSurf)
	// Custom middleware to load and save session data for the request
	mux.Use(SessionLoad)

	// Define routes and their corresponding handler functions
	// Route for the home page
	mux.Get("/", handlers.Repo.Home)
	// Route for the about page
	mux.Get("/about", handlers.Repo.About)
	// Route for Generals Quarters room page
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	// Route for Majors Suite room page
	mux.Get("/majors-suite", handlers.Repo.Majors)
	// Route for the contact page
	mux.Get("/contact", handlers.Repo.Contact)
	// Route for making a reservation
	mux.Get("/reservation", handlers.Repo.Reservation)
	// Route for searching availability
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	// Route for making a reservation
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	// Serve static files (like CSS, JS, images) from the "./static" directory
	// Create a file server for the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	// Handle requests for static files
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Return the configured router
	return mux
}
