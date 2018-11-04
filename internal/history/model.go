package history

// LogEntry models the resuls of a particular drawing
type LogEntry struct {
	ID      int    `json:"id"`
	Barista string `json:"barista"`
	Cleaner string `json:"cleaner"`
	DrawnAt string `json:"drawnAt"`
}
