package renderer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders a page a templage using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl,
		"./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderCachedTemplate(w http.ResponseWriter, templateName string) {

	var result *template.Template
	var err error

	//check to see if we already have a cached template
	_, inMap := templateCache[templateName]

	if !inMap {

		// here we go with a new template
		log.Println("caching the created template")
		err = createTemplateCache(templateName)
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Println("using cached template")
	}

	result = templateCache[templateName]
	err = result.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(templateName string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	templateCache[templateName] = tmpl
	return nil
}
