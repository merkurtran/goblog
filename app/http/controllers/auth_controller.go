package controllers

import (
	"net/http"

	"github.com/merkurtran/goblog/pkg/view"
)

type AuthController struct{}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "auth.Register")
}

func (*AuthController) DoRegister(w http.Response, r *http.Request) {
	//
}
