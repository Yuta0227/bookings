package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yuta0227/bookings/pkg/config"
	"github.com/Yuta0227/bookings/pkg/handlers"
	"github.com/Yuta0227/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = "localhost:8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Error creating template cache", err)
		return
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
	fmt.Printf("Starting application on port %s\n", portNumber)
}
