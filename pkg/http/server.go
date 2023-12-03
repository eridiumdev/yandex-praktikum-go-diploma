package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		Router *gin.Engine
		Server *nethttp.Server
	}
	ServerConfig struct {
		Addr string
	}
)

func NewServer(cfg ServerConfig, middlewares ...gin.HandlerFunc) *Server {
	handler := gin.New()
	handler.ContextWithFallback = true

	for _, m := range middlewares {
		handler.Use(m)
	}
	//handler.Handler()

	return &Server{
		Router: handler,
		Server: &nethttp.Server{
			Handler: handler,
			//ReadTimeout:  60 * time.Second,
			//WriteTimeout: 60 * time.Second,
			Addr: cfg.Addr,
		},
	}
}
