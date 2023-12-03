package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eridiumdev/yandex-praktikum-go-diploma/internal/usecase"
	pkghttp "github.com/eridiumdev/yandex-praktikum-go-diploma/pkg/http"
	"github.com/eridiumdev/yandex-praktikum-go-diploma/pkg/logger"
)

type ShortenerController struct {
	shortener usecase.Shortener

	log *logger.Logger
}

func NewShortenerController(srv *pkghttp.Server, shortener usecase.Shortener, log *logger.Logger) *ShortenerController {
	c := &ShortenerController{
		shortener: shortener,
		log:       log,
	}

	srv.Router.GET("/ping", c.ping)

	return c
}

func (ct *ShortenerController) ping(c *gin.Context) {
	ctx := c

	err := ct.shortener.Ping(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
