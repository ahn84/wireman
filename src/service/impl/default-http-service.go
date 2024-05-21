package impl

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	// docs "go-app/docs"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"

	"wgd/src/configuration"
	"wgd/src/service/iface"
	"wgd/src/service/rest"

	oidc "github.com/TJM/gin-gonic-oidcauth"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @host      localhost:8866
// @BasePath  /api/v2

type httpService struct {
	server    *http.Server
	engine    *gin.Engine
	wgManager iface.WireguardSPI
}

func NewHttpService(wgm iface.WireguardSPI) *httpService {
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
		server:    server,
		engine:    engine,
		wgManager: wgm,
	}
}

// Initialize setup routes and handlers
func (s *httpService) Initialize(context.Context) error {
	// ctx.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// docs.SwaggerInfo.BasePath = "/api/v2"
	// apiGroup := ctx.R.Group("/api/v2")

	store := cookie.NewStore([]byte("secret"), nil)    // Do not use "secret", nil in production. This sets the keypairs for auth, encryption of the cookies.
	s.engine.Use(sessions.Sessions("oidcauth", store)) // Sessions must be Use(d) before oidcauth, as oidcauth requires sessions

	wgHandler := rest.NewWireguardHandler(s.wgManager)
	wgGroup := s.engine.Group("/api/wg")
	wgGroup.GET("/add-peers", wgHandler.AddPeers)

	authGroup := s.engine.Group("/api/auth")
	oidcConfig := oidc.Config{
		IssuerURL:    "https://id.lab.linksafe.vn/realms/ztwg",
		ClientID:     "localdev",
		Scopes:       []string{"openid", "profile", "email"},
		ClientSecret: "dXVgPdCoVrXQEH2f4iqMCre8dGi1uc6J",
		RedirectURL:  "http://localhost:8080/api/auth/callback",
		LogoutURL:    "http://localhost:8080/api/auth/logout",
	}
	oidcAuth, _ := oidc.GetOidcAuth(&oidcConfig)
	authGroup.GET("/login", oidcAuth.Login) // Unnecessary, as requesting a "AuthRequired" resource will initiate login, but potentially convenient
	authGroup.GET("/callback", oidcAuth.AuthCallback)
	authGroup.GET("/logout", oidcAuth.Logout)

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
