package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/merkurtran/goblog/app/models/article"
	"github.com/merkurtran/goblog/pkg/logger"
	"github.com/merkurtran/goblog/pkg/route"
	"github.com/merkurtran/goblog/pkg/types"
	"gorm.io/gorm"
)

type ArticlesController struct{}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 article not found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server error")
		}
	} else {
		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{"RouteName2URL": route.Name2URL, "Uint64ToString": types.Uint64ToString}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)
		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
	fmt.Fprint(w, "articles ID："+id)
}
