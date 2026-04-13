package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/merkurtran/goblog/pkg/logger"
	"github.com/merkurtran/goblog/pkg/route"
	"github.com/merkurtran/goblog/pkg/types"
)

type PagesController struct{}

func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, Welcome to goblog！</h1>")
}

func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,
		"This blog is intended for recording programming notes. If you have any feedback or suggestions, please contact us. "+
			"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>The requested page was not found :(</h1><p>If you have any questions, please contact us.</p>")
}

type ArticlesController struct{}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 article not found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server error")
		}
	} else {
		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{"RouteName2URL": route.Name2URL, "Int64ToString": types.Int64ToString}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)
		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
	fmt.Fprint(w, "articles ID："+id)
}
