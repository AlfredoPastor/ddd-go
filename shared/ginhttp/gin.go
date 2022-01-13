package ginhttp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AlfredoPastor/ddd-go/shared/ginhttp/middleware/logging"
	"github.com/AlfredoPastor/ddd-go/shared/ginhttp/middleware/recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config interface {
	GetHTTPEnv() *HTTPEnv
}

//HTTPEnv obtiene la configuracion desde el entorno
type HTTPEnv struct {
	HTTPAddr string `env:"HTTP_ADDR"`
	HTTPPort int    `env:"HTTP_PORT"`
}

type HttpServer struct {
	*http.Server
	*gin.Engine
}

func (h HttpServer) Run(ctx context.Context) error {
	log.Println("Server running on:", h.Server.Addr)

	go func() {
		if err := h.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return h.Server.Shutdown(ctxShutDown)
}

func NewHttpServer(config Config) HttpServer {
	engine := gin.New()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.GetHTTPEnv().HTTPAddr, config.GetHTTPEnv().HTTPPort),
		Handler: engine,
	}
	engine.Use(recovery.Middleware(), logging.Middleware(), cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	return HttpServer{Server: srv, Engine: engine}
}
