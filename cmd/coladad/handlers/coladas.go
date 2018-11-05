package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/apex/log"
	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/config"
	"github.com/rippinrobr/lunch-n-learn/internal/activities"
	"github.com/rippinrobr/lunch-n-learn/internal/db"
	"github.com/rippinrobr/lunch-n-learn/internal/drawing"
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

func (d *Drinker) GetBaristaAndCleaner(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	lastDraw, err := db.GetPreviousResult(d.DB)
	if err != nil {
		if !strings.Contains(err.Error(), "no rows in result set") {
			log.Warn("Unable to find previous drawing results")
			web.Error(ctx, w, err)
			return nil
		}

		lastDraw = &drawing.Result{}
	}

	drinkers, err := db.GetDrinkers(d.DB)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			log.Warn("Unable to find drinkers from drinkers table")
			web.Error(ctx, w, web.ErrNotFound)
			return nil
		}
	}
	brew := activities.Brew{}
	barista := brew.PickNextBarista(drinkers, lastDraw)

	clean := activities.Clean{}
	cleaner := clean.PickNextCleaner(drinkers, barista, lastDraw)

	results := drawing.CreateNewResult(barista, cleaner)
	results, err = db.AddResult(d.DB, results)
	if err != nil {
		log.Warn("Unable to add a result to the drawings table: " + err.Error())
		web.Error(ctx, w, err)
		return nil
	}

	web.Respond(ctx, w, results, http.StatusOK)
	return nil
}
