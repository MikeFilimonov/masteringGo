package handlers

import (
	"net/http"

	"github.com/MikeFilimonov/masteringGo/pkg/renderer"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "about.page.tmpl")
}
