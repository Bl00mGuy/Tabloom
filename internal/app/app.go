package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bl00mGuy/Tabloom/internal/config"
	"github.com/Bl00mGuy/Tabloom/pkg/logger"
)

// App represents the main application
type App struct {
	config *config.Config // Application configuration
	server *http.Server   // HTTP server
}

// New creates a new application instance
func New(version, commit string) (*App, error) {
	// Display application banner
	printBanner(version)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Initialize logger
	if err := logger.Init(cfg.Log.Level); err != nil {
		return nil, err
	}

	log.Printf("Tabloom v%s starting on port %s", version, cfg.Server.Port)

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      initRoutes(), // TODO: initialize routes
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &App{
		config: cfg,
		server: server,
	}, nil
}

// Run starts the application
func (a *App) Run() error {
	// Start server in goroutine
	go func() {
		log.Printf("Server running on http://localhost:%s", a.config.Server.Port)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}

// printBanner displays ASCII application banner
func printBanner(version string) {
	banner := `
████████╗█████╗ ██████╗ ██╗      ██████╗  ██████╗ ███╗   ███╗
╚══██╔══╝██╔══██╗██╔══██╗██║     ██╔═══██╗██╔═══██╗████╗ ████║
   ██║   ███████║██████╔╝██║     ██║   ██║██║   ██║██╔████╔██║
   ██║   ██╔══██║██╔══██╗██║     ██║   ██║██║   ██║██║╚██╔╝██║
   ██║   ██║  ██║██████╔╝███████╗╚██████╔╝╚██████╔╝██║ ╚═╝ ██║
   ╚═╝   ╚═╝  ╚═╝╚═════╝ ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝     ╚═╝
                                                              
                Musical Tab Workshop v%s
`
	log.Printf(banner, version)
}

// initRoutes initializes application routes
func initRoutes() http.Handler {
	mux := http.NewServeMux()

	// Basic routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)

	return mux
}

// homeHandler handles the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Welcome to Tabloom! 🎸\n\nString instrument tab workshop"))
}

// healthHandler handles health checks
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok", "service": "tabloom"}`))
}
