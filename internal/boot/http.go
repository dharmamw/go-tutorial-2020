package boot

import (
	"log"
	"net/http"

	"tugas-arif/internal/config"

	"github.com/jmoiron/sqlx"

	arifData "tugas-arif/internal/data/arif"
	arifServer "tugas-arif/internal/delivery/http"
	arifHandler "tugas-arif/internal/delivery/http/arif"
	arifService "tugas-arif/internal/service/arif"

	firebaseclient "tugas-arif/pkg/firebaseClient"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s   arifServer.Server    // HTTP Server Object
		ad  arifData.Data        // User domain data layer
		as  arifService.Service  // User domain service layer
		ah  *arifHandler.Handler // User domain handler
		cfg *config.Config       // Configuration object
		fb  *firebaseclient.Client
	)

	// Get configuration
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()

	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	fb, err = firebaseclient.NewClient(cfg)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// User domain initialization
	ad = arifData.New(db, fb)
	as = arifService.New(ad)
	ah = arifHandler.New(as)

	// Inject service used on handler

	s = arifServer.Server{
		Arif: ah,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
