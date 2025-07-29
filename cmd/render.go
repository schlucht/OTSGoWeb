package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
}

var defaultPartials = []string{"js"} //[2]string{"header", "navigation"}
var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {

	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	partials = append(partials, defaultPartials...)
	templateToRender := fmt.Sprintf("templates/pages/%s.page.tmpl", page)

	_, templateInMap := app.templateCache[templateToRender]

	if templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, r)
	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/partials/%s.partial.tmpl", x)
		}
	}

	// read Layouts
	layouts, err := app.parseLayouts()
	if err != nil {
		app.errorLog.Println("Find no layouts", err)
		return nil, err
	}

	partials = append(partials, templateToRender)
	partials = append(partials, layouts...)

	t, err = template.New(fmt.Sprintf("%s.page.tmpl", page)).Funcs(functions).ParseFS(templateFS, partials...)
	if err != nil {
		app.errorLog.Println("Error parsing template files", err)
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}

func (app *application) parseLayouts() ([]string, error) {
	var paths []string
	matches, err := filepath.Glob("./cmd/templates/layouts/*")
	if err != nil {
		return nil, err
	}
	for _, path := range matches {
		_, path = filepath.Split(path)
		paths = append(paths, filepath.Join("templates/layouts", path))
	}
	return paths, nil
}
