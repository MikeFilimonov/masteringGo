package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MikeFilimonov/masteringGo/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders a page a templage using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache

	templateToRender, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buffer := new(bytes.Buffer)

	_ = templateToRender.Execute(buffer, nil)

	_, err := buffer.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

	// // render the template
	// _, err = buffer.WriteTo(w)

	// if err != nil {
	// 	log.Println(err)
	// }

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	// get all of the files named *.page.tml from .templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {

		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {

			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = templateSet

	}

	return myCache, nil

}
