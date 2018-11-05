package history

// DrawingResult models the resuls of a particular drawing
type DrawingResult struct {
	ID         int    `json:"id"`
	Barista    string `json:"barista"`
	BaristaID  int    `json:"baristaId"`
	BaristaImg string `json:"baristaImg"`
	Cleaner    string `json:"cleaner"`
	CleanerID  int    `json:"cleanerId"`
	CleanerImg string `json:"cleanerImg"`
	DrawnAt    string `json:"drawnAt"`
}
