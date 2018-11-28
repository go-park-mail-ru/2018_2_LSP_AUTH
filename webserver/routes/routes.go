package routes

import (
	"net/http"

	"github.com/go-park-mail-ru/2018_2_LSP_AUTH/webserver/handlers"
	"github.com/go-park-mail-ru/2018_2_LSP_AUTH/webserver/middlewares"
)

func Get() handlers.HandlersMap {
	handlersMap := handlers.HandlersMap{}
	handlersMap["/"] = makeRequest(handlers.HandlersMap{
		"post":   handlers.AuthUserHandler,
		"delete": handlers.ClearCookiesHandler,
	})
	return handlersMap
}

func makeRequest(handlersMap handlers.HandlersMap) handlers.HandlerFunc {
	return func(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
		var key string
		switch r.Method {
		case http.MethodGet:
			key = "get"
		case http.MethodPost:
			key = "post"
		case http.MethodPut:
			key = "put"
		case http.MethodDelete:
			key = "delete"
		default:
			return middlewares.Cors(handlers.DefaultHandler)(env, w, r)
		}
		if _, ok := handlersMap[key]; ok {
			return middlewares.Cors(handlersMap[key])(env, w, r)
		}
		return middlewares.Cors(handlers.DefaultHandler)(env, w, r)
	}
}
