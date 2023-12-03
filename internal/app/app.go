package app

import (
	"context"
	"errors"
	nethttp "net/http"

	"github.com/gin-gonic/gin"

	"github.com/eridiumdev/yandex-praktikum-go-diploma/config"
	"github.com/eridiumdev/yandex-praktikum-go-diploma/internal/controller/http"
	"github.com/eridiumdev/yandex-praktikum-go-diploma/internal/usecase"
	pkghttp "github.com/eridiumdev/yandex-praktikum-go-diploma/pkg/http"
	"github.com/eridiumdev/yandex-praktikum-go-diploma/pkg/http/middleware"
	"github.com/eridiumdev/yandex-praktikum-go-diploma/pkg/logger"
)

type App struct {
	server *pkghttp.Server
	log    *logger.Logger
}

func New(ctx context.Context, cfg *config.Config, log *logger.Logger) (*App, error) {

	server := pkghttp.NewServer(cfg.Server,
		gin.Recovery(),
		middleware.RequestID(log),
		middleware.Logger(log.SubLogger("http_requests")))

	shortenerUC := usecase.NewShortener(log.SubLogger("shortener_uc"))

	http.NewShortenerController(server, shortenerUC, log.SubLogger("shortener_controller"))

	return &App{
		server: server,
		log:    log,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.log.Info(ctx).Msgf("Listening on %s", a.server.Server.Addr)
	if err := a.server.Server.ListenAndServe(); err != nil && !errors.Is(err, nethttp.ErrServerClosed) {
		return err
	}
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	err := a.server.Server.Shutdown(ctx)
	if err != nil {
		return a.log.Wrap(err, "shutdown server")
	}

	return nil
}
