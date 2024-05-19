package impl

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"wgd/src/configuration"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type httpService struct {
	server *http.Server
	engine *gin.Engine
}

func NewHttpService() *httpService {
	engine := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	engine.Use(cors.New(corsConfig)) //enable cors default to all hosts

	// engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// appContext.R = engine
	// Setup routing
	engine.GET("/healthz", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte(viper.GetString(configuration.CFG_APPLICATION_NAME)))
	})

	server := &http.Server{
		Addr:           viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &httpService{
		server: server,
		engine: engine,
	}
}

func (s *httpService) Initialize(context.Context) error {
	return nil
}

func (s *httpService) Start(context.Context) error {
	go func() {
		slog.Info("Starting HTTP server", "addr", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Warn("ListenAndServe", "error", err)
		}
	}()
	return nil
}
func (s *httpService) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *httpService) HealthCheck(context.Context) error {
	return nil
}

func (s *httpService) GetServer() *http.Server {
	return s.server
}

func (s *httpService) GetEngine() *gin.Engine {
	return s.engine
}
