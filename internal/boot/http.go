package boot

import (
	"log"
	"net/http"

	"go-tutorial-2020/internal/config"

	userData "go-tutorial-2020/internal/data/user"
	server "go-tutorial-2020/internal/delivery/http"
	userHandler "go-tutorial-2020/internal/delivery/http/user"
	userService "go-tutorial-2020/internal/service/user"
	firebaseclient "go-tutorial-2020/pkg/firebaseClient"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s   server.Server        // HTTP Server Object
		ud  userData.Data        // User domain data layer
		us  userService.Service  // User domain service layer
		uh  *userHandler.Handler // User domain handler
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
	// db, err := sqlx.Open("mysql", cfg.Database.Master)
	// if err != nil {
	// 	log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	// }
	// Open firebase DB Connection
	fb, err = firebaseclient.NewClient(cfg)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize firebase connection: %v", err)
	}

	// User domain initialization
	ud = userData.New(fb)
	us = userService.New(ud)
	uh = userHandler.New(us)

	// Inject service used on handler
	s = server.Server{
		User: uh,
	}

	// Error Handling
	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
