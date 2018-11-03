package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/config"
	"github.com/rippinrobr/lunch-n-learn/internal/db"
	"github.com/rippinrobr/lunch-n-learn/internal/platform/web"
)

// Client represents the Client API method handler set.
type Drinker struct {
	DB  *sql.DB
	cfg config.Config
}

func (d *Drinker) List(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	drinkers, err := db.GetDrinkers(d.DB)
	if err != nil {
		web.Error(ctx, w, err)
		return err
	}

	web.Respond(ctx, w, drinkers, http.StatusOK)
	return nil
}
