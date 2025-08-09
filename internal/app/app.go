package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gsvd/website/internal/blog"
	"github.com/gsvd/website/internal/handlers"
	"github.com/gsvd/website/internal/tmplengine"
	"github.com/gsvd/website/web"
)

var (
	ErrBadEnvConfiguration = fmt.Errorf("bad environment configuration please set: ENV, HOST and PORT")
)

type App struct {
	Logger         *slog.Logger
	TemplateEngine *tmplengine.TmplEngine
	Articles       map[string]*blog.Article
	Router         chi.Router
	Addr           string
}

func New(commitHash string) (*App, error) {
	if err := validateEnvConfiguration(); err != nil {
		return nil, fmt.Errorf("environment validation failed: %w", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	templateEngine, err := tmplengine.New(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed to init template engine: %w", err)
	}

	return &App{
		Logger:         logger,
		TemplateEngine: templateEngine,
		Router:         r,
		Addr:           fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
	}, nil
}

func (a *App) RegisterRoutes() {
	vh := handlers.NewViewHandler(a.TemplateEngine, a.Logger)

	a.Router.Get("/", vh.ShowIndex)
	a.Router.Get("/contact", vh.ShowContact)
	a.Router.Get("/resume", vh.ShowResume)
	a.Router.Get("/blog/{slug:[a-z-]+}", vh.ShowArticle)
	a.Router.Get("/blog", vh.ShowBlog)

	// Static file handler
	a.Router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(web.StaticFS()))))
	a.logRegisteredRoutes()
}

func (a *App) logRegisteredRoutes() {
	routes := a.Router.Routes()
	for _, route := range routes {
		for method := range route.Handlers {
			a.Logger.Debug("Registered route",
				"method", method,
				"path", route.Pattern,
			)
		}
	}
}

func (a *App) Start() error {
	a.Logger.Info("Starting server on", "addr", a.Addr)
	return http.ListenAndServe(a.Addr, a.Router)
}

func validateEnvConfiguration() error {
	if os.Getenv("HOST") == "" || os.Getenv("PORT") == "" || os.Getenv("ENV") == "" {
		return ErrBadEnvConfiguration
	}

	return nil
}
