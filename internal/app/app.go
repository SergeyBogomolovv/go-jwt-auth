package app

import (
	"context"
	"go-jwt-auth/internal/config"
	"go-jwt-auth/internal/controllers"
	"go-jwt-auth/internal/middleware"
	"go-jwt-auth/internal/repositories"
	"go-jwt-auth/internal/usecases"
	"log/slog"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type App struct {
	db     *sqlx.DB
	router *http.ServeMux
	cfg    *config.Config
}

func New(db *sqlx.DB, cfg *config.Config) *App {
	return &App{
		db:     db,
		router: http.NewServeMux(),
		cfg:    cfg,
	}
}

func (app *App) RegisterRoutes() {
	validate := validator.New(validator.WithRequiredStructEnabled())

	usersRepository := repositories.NewUserRepository(app.db)

	authUsecase := usecases.NewAuthUsecase(usersRepository, app.cfg.JwtSecret)
	usersUsecase := usecases.NewUsersUsecase(usersRepository)

	authMiddleware := middleware.NewAuthMiddleware(authUsecase)
	controllers.NewAuthController(authUsecase, validate).RegisterRoutes(app.router)
	controllers.NewUsersController(usersUsecase, validate).RegisterRoutes(app.router, authMiddleware.Middleware)
}

func (app *App) Run() {
	addr := app.cfg.Host + ":" + strconv.Itoa(int(app.cfg.Port))

	app.RegisterRoutes()
	server := &http.Server{
		Addr:    addr,
		Handler: middleware.Logger(app.router),
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
