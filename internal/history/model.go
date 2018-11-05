package history

// LogEntry models the resuls of a particular drawing
type LogEntry struct {
	ID         int    `json:"id"`
	Barista    string `json:"barista"`
	BaristaID  int    `json:"baristaId"`
	BaristaImg string `json:"baristaImg"`
	Cleaner    string `json:"cleaner"`
	CleanerID  int    `json:"cleanerId"`
	CleanerImg string `json:"baristaImg"`
	DrawnAt    string `json:"drawnAt"`
}
