package middlewares

import (
	"net/http"
	"strings"

	"github.com/go-park-mail-ru/2018_2_LSP_AUTH/webserver/handlers"
)

// Cors Middleware that enables CORS
func Cors(next handlers.HandlerFunc) handlers.HandlerFunc {
	return func(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
		if origin := r.Header.Get("Origin"); origin != "" {
			originSplitted := strings.Split(origin, ":")
			if strings.HasSuffix(originSplitted[0], ".jackal.online") {
				if len(originSplitted) > 1 {
					w.Header().Set("Access-Control-Allow-Origin", originSplitted[0]+":"+originSplitted[1])
				} else {
					w.Header().Set("Access-Control-Allow-Origin", originSplitted[0])
				}
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "jackal.online")
			}
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		if r.Method == http.MethodOptions {
			return nil
		}

		w.Header().Set("Content-Type", "application/json")

		return next(env, w, r)
	}
}
