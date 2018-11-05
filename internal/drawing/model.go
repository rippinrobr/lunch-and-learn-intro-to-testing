package drawing

import (
	"time"

	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
)

// Result models the resuls of a particular drawing
type Result struct {
	ID         int64  `json:"id"`
	Barista    string `json:"barista"`
	BaristaID  int64  `json:"baristaId"`
	BaristaImg string `json:"baristaImg"`
	Cleaner    string `json:"cleaner"`
	CleanerID  int64  `json:"cleanerId"`
	CleanerImg string `json:"cleanerImg"`
	DrawnAt    string `json:"drawnAt"`
}

// CreateNewResult takes the newly chosen barista and cleaner to create a Result
// record
func CreateNewResult(barista *coladas.Drinker, cleaner *coladas.Drinker) *Result {
	return &Result{
		Barista:    barista.Name,
		BaristaID:  barista.UID,
		BaristaImg: barista.HeadshotPath,
		Cleaner:    cleaner.Name,
		CleanerID:  cleaner.UID,
		CleanerImg: cleaner.HeadshotPath,
		DrawnAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
}
