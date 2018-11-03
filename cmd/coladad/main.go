package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/config"
	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/handlers"
)

// TODO: Create base resource /coladas
// GET routes /coladas/drinker  /coladas/cleaner
// POST route /coladas/brew     /coladas/clean
func main() {

	// ============================================================
	// Get Config
	cfg := config.New()

	// ============================================================
	// get db connection
	db, err := sql.Open(cfg.DBType, cfg.DBConnInfo)
	if err != nil {
		log.Fatalf("startup : Register DB : %v", err)
	}

	// ============================================================
	// Start Service
	server := http.Server{
		Addr:           cfg.APIHost,
		Handler:        handlers.API(db, cfg),
		MaxHeaderBytes: 1 << 20,
	}

	// Starting the service, listening for requests.
	go func() {
		log.Printf("startup : Listening %s", cfg.APIHost)
		log.Printf("shutdown : Listener closed : %v", server.ListenAndServe())
	}()

	// ============================================================
	// Shutdown
	// Blocking main and waiting for shutdown.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	<-osSignals

	log.Println("main : Completed")
}
