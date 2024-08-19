package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/justinas/nosurf"
	"github.com/Yuta0227/bookings/pkg/config"
	"github.com/Yuta0227/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
    td.CSRFToken = nosurf.Token(r)
    return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
    var templateCache map[string]*template.Template
    if app.UseCache {
        templateCache = app.TemplateCache
    } else {
        templateCache, _ = CreateTemplateCache()
    }

    template, ok := templateCache[tmpl]
    if !ok {
        log.Println("Could not get template from template cache")
        return
    }

    buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	_ = template.Execute(buf, td)
	// render the template
	_, err := buf.WriteTo(w)
    if err != nil {
		log.Println("Error writing template to browser", err)
    }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		// extract the file name (like home.page.tmpl)
		name := filepath.Base(page)
		// parse the page template file in to a set of templates
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// check if there is a layout template
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// if there are any layout templates, parse them
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// add the template set to the cache
		myCache[name] = templateSet
	}
	// return the cache
	return myCache, nil
}
