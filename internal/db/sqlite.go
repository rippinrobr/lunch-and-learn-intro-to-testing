package db

import (
	"database/sql"
	"errors"

	"github.com/apex/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
	"github.com/rippinrobr/lunch-n-learn/internal/drawing"
)

// GetDrinkers returns a slice of coladas.Drinker structs
func GetDrinkers(db *sql.DB) ([]*coladas.Drinker, error) {
	if db == nil {
		return nil, errors.New("[GetDrinkers] db was nil")
	}

	rows, err := db.Query("SELECT * FROM drinkers order by name")
	if err != nil {
		return nil, err
	}

	drinkers := make([]*coladas.Drinker, 0)
	var uid int64
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
		drinkers = append(drinkers, d)
	}

	return drinkers, nil
}

// GetPreviousResult returns the latest Colada Lottery drawing
// results
func GetPreviousResult(db *sql.DB) (*drawing.Result, error) {
	if db == nil {
		return nil, errors.New("[GetDrinkers] db was nil")
	}

	var id int64
	var barista string
	var baristaID int64
	var baristaImg string
	var cleaner string
	var cleanerID int64
	var cleanerImg string
	var drawnAt string

	row := db.QueryRow("SELECT id, barista, baristaId, baristaImg, cleaner, cleanerId, cleanerImg, drawn_at FROM drawings order by id desc limit 1")
	err := row.Scan(&id, &barista, &baristaID, &baristaImg, &cleaner, &cleanerID, &cleanerImg, &drawnAt)
	if err != nil {
		return nil, err
	}

	return &drawing.Result{
		ID:         id,
		Barista:    barista,
		BaristaID:  baristaID,
		BaristaImg: baristaImg,
		Cleaner:    cleaner,
		CleanerID:  cleanerID,
		CleanerImg: cleanerImg,
		DrawnAt:    drawnAt,
	}, nil
}

// AddResult adds the given drawing.Result into the drawings table
func AddResult(db *sql.DB, res *drawing.Result) (*drawing.Result, error) {
	if db == nil {
		return nil, errors.New("[AddResult] db was nil")
	}

	if res == nil {
		return nil, errors.New("[AddResult] cannot add a drawing result that is nil")
	}

	result, err := db.Exec("INSERT INTO drawings (barista, baristaId, baristaImg, cleaner, cleanerId, cleanerImg, drawn_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		res.Barista, res.BaristaID, res.BaristaImg, res.Cleaner, res.CleanerID, res.CleanerImg, res.DrawnAt)
	if err != nil {
		return nil, err
	}

	res.ID, err = result.LastInsertId()
	if err != nil {
		return nil, errors.New("[AddResult] No record added: " + err.Error())
	}

	return res, nil
}
