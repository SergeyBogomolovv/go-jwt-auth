package app

import (
	"context"
	"go-jwt-auth/internal/config"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

type App struct {
	db     *sqlx.DB
	router *chi.Mux
	cfg    *config.Config
}

func New(db *sqlx.DB, cfg *config.Config) *App {
	return &App{
		db:     db,
		router: chi.NewRouter(),
		cfg:    cfg,
	}
}

func (app *App) RegisterRoutes() {
	app.router.Use(middleware.RequestID)
	app.router.Use(middleware.RealIP)
	app.router.Use(middleware.Logger)
	app.router.Use(middleware.Recoverer)

	app.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

func (app *App) Run() {
	addr := app.cfg.Host + ":" + strconv.Itoa(int(app.cfg.Port))

	app.RegisterRoutes()
	server := &http.Server{
		Addr:    addr,
		Handler: app.router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Server started on %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server stopped")
}
