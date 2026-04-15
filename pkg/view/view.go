package view

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/merkurtran/goblog/pkg/logger"
	"github.com/merkurtran/goblog/pkg/route"
)

type D map[string]interface{}

func Render(w io.Writer, data interface{}, tplFiles ...string) {
	viewDir := "resources/views/"
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	allFiles := append(layoutFiles, tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}

func RenderSimple(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data interface{}, tplFiles ...string) {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	allFiles := append(layoutFiles, tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}
