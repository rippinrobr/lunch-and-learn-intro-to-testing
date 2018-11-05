package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/config"
	"github.com/rippinrobr/lunch-n-learn/internal/activities"
	"github.com/rippinrobr/lunch-n-learn/internal/db"
	"github.com/rippinrobr/lunch-n-learn/internal/history"
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
	lastDraw, err := db.GetPreviousDrawResult(d.DB)
	if err != nil {
		if !strings.Contains(err.Error(), "no rows in result set") {
			log.Warn("Unable to find previous drawing results")
			web.Error(ctx, w, err)
			return nil
		}

		lastDraw = &history.LogEntry{}
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

	le := history.LogEntry{
		ID:         1,
		Barista:    barista.Name,
		BaristaID:  barista.UID,
		BaristaImg: barista.HeadshotPath + " image",
		Cleaner:    cleaner.Name,
		CleanerID:  cleaner.UID,
		CleanerImg: cleaner.HeadshotPath,
		DrawnAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	fmt.Printf("le: %+v\n", le)
	web.Respond(ctx, w, le, http.StatusOK)
	return nil
}
