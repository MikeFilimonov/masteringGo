package renderer

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a page a templage using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// create a template cache
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache

	templateToRender, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	err = templateToRender.Execute(buffer, nil)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buffer.WriteTo(w)

	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {

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
