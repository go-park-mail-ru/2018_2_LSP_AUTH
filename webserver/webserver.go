package webserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2018_2_LSP_AUTH/webserver/handlers"
	"github.com/go-park-mail-ru/2018_2_LSP_AUTH/webserver/routes"

	zap "go.uber.org/zap"
)

// Run Run webserver on specified port (passed as string the
// way regular http.ListenAndServe works)
func Run(addr string, db *sql.DB) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	env := &handlers.Env{
		DB:     db,
		Logger: sugar,
	}

	handlersMap := routes.Get()
	for URL, h := range handlersMap {
		http.Handle(URL, handlers.Handler{env, h})
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}
