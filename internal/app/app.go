package app

import (
	"context"
	"go-jwt-auth/internal/config"
	"go-jwt-auth/internal/controllers"
	mw "go-jwt-auth/internal/middleware"
	"go-jwt-auth/internal/repositories"
	"go-jwt-auth/internal/usecases"
	"log/slog"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New(validator.WithRequiredStructEnabled())

	usersRepository := repositories.NewUserRepository(app.db)

	authUsecase := usecases.NewAuthUsecase(usersRepository, app.cfg.JwtSecret)
	usersUsecase := usecases.NewUsersUsecase(usersRepository)

	authMiddleware := mw.NewAuthMiddleware(authUsecase)
	controllers.NewAuthController(authUsecase, validate).RegisterRoutes(app.router)
	controllers.NewUsersController(usersUsecase, validate).RegisterRoutes(app.router, authMiddleware.Middleware)
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
		slog.Info("Starting server", "addr", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Error starting server", "error", err)
		}
	}()
	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Error shutting down server", "error", err)
	}

	slog.Info("Server stopped")
}
