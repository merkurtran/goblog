package controllers

import (
	"fmt"
	"net/http"
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
