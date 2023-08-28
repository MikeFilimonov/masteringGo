package handlers

import (
	"net/http"

	"github.com/MikeFilimonov/masteringGo/pkg/config"
	"github.com/MikeFilimonov/masteringGo/pkg/renderer"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(appConfig *config.AppConfig) *Repository {

	return &Repository{
		App: appConfig,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//renderer.RenderTemplate(w, "home.page.tmpl")
	renderer.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	//renderer.RenderTemplate(w, "about.page.tmpl")
	renderer.RenderTemplate(w, "about.page.tmpl")
}
