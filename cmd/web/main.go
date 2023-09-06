package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MikeFilimonov/masteringGo/pkg/config"
	"github.com/MikeFilimonov/masteringGo/pkg/handlers"
	"github.com/MikeFilimonov/masteringGo/pkg/renderer"
)

const portNumber = ":8080"

// main is the main entry point of the app
func main() {

	var app config.AppConfig

	templateCache, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renderer.NewTemplates(&app)
	fmt.Printf("Starting the app on port %s ", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
