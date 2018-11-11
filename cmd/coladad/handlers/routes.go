package handlers

import (
	"database/sql"
	"net/http"

	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/config"
	"github.com/rippinrobr/lunch-n-learn/internal/middleware"
	"github.com/rippinrobr/lunch-n-learn/internal/platform/web"
)

func API(coladaDB *sql.DB, cfg config.Config) http.Handler {

	// Create the application.
	app := web.New(middleware.RequestLogger, middleware.ErrorHandler)

	// ============================================================
	// Colada Drinkers API
	d := Drinker{
		DB:  coladaDB,
		cfg: cfg,
	}
	app.Handle("GET", "/v1/drinkers", d.List)

	// ============================================================
	// The Lottery Drawing API
	h := Drawing{
		DB:  coladaDB,
		cfg: cfg,
	}
	app.Handle("GET", "/v1/drawings/previous", h.GetLatest)
	app.Handle("GET", "/v1/drawings", d.GetBaristaAndCleaner)

	return app
}
