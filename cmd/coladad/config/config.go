package config

// Config houses the informaiton
// needed for the Colada api
type Config struct {
	APIHost    string
	DBConnInfo string
	DBType     string
}

// New creates a new Config object
func New() Config {
	return Config{
		APIHost:    ":9999",
		DBConnInfo: "/Users/robertrowe/go/src/github.com/rippinrobr/lunch-n-learn/database/colada-lottery.db",
		DBType:     "sqlite3",
	}
}
