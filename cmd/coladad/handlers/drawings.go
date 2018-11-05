package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/config"
	"github.com/rippinrobr/lunch-n-learn/internal/db"
	"github.com/rippinrobr/lunch-n-learn/internal/platform/web"
)

// Drawing represents the History API method handler set.
type Drawing struct {
	DB  *sql.DB
	cfg config.Config
}

// GetLatest gets the most recent drawing results
func (d *Drawing) GetLatest(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	lastDraw, err := db.GetPreviousResult(d.DB)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			log.Println("[GetLatest] No drawing found")
			web.Error(ctx, w, web.ErrNotFound)
			return nil
		}

		web.Error(ctx, w, err)
		return nil
	}

	web.Respond(ctx, w, lastDraw, http.StatusOK)
	return nil
}
