package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}

func Name2URL(routeName string, pairs ...string) string {
	r := route.Get(routeName)
	if r == nil {
		return ""
	}
	url, err := r.URL(pairs...)
	if err != nil {
		return ""
	}
	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
