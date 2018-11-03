package history

type LogEntry struct {
	ID      int    `json:"id"`
	Barista string `json:"barista"`
	Cleaner string `json:"cleaner"`
}
