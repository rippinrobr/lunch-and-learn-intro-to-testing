package db

import (
	"database/sql"
	"fmt"

	"github.com/apex/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
	"github.com/rippinrobr/lunch-n-learn/internal/history"
)

func GetDrinkers(db *sql.DB) ([]coladas.Drinker, error) {
	rows, err := db.Query("SELECT * FROM drinkers order by name")
	if err != nil {
		return nil, err
	}

	drinkers := make([]coladas.Drinker, 0)
	var uid int
	var name string
	var canMake int
	var headshotPath string
	for rows.Next() {
		err = rows.Scan(&uid, &name, &canMake, &headshotPath)
		if err != nil {
			return nil, err
		}

		d, err := coladas.CreateColadaDrinker(uid, name, canMake == 1, headshotPath)
		if err != nil {
			log.Infof("Unable to create a Drinker for uid: %d\n", uid)
			continue
		}
		drinkers = append(drinkers, *d)
	}

	return drinkers, nil
}

func GetPreviousDrawResult(db *sql.DB) (*history.LogEntry, error) {
	var id int
	var barista string
	var cleaner string
	var drawnAt string

	row := db.QueryRow("SELECT id, barista, cleaner, drawn_at FROM history order by id desc limit 1")
	err := row.Scan(&id, &barista, &cleaner, &drawnAt)
	if err != nil {
		return nil, err
	}

	fmt.Printf("id: '%d' barista: '%s'", id, barista)
	return &history.LogEntry{
		ID:      id,
		Barista: barista,
		Cleaner: cleaner,
		DrawnAt: drawnAt,
	}, nil
}
