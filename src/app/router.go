package app

import (
	// docs "go-app/docs"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"

	"wgd/src/configuration"
	"wgd/src/service/iface"
	"wgd/src/service/rest"
)

// @host      localhost:8866
// @BasePath  /api/v2

// NewRouter defines api paths
func InitRouting(ctx *configuration.AppContext) {
	// ctx.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// docs.SwaggerInfo.BasePath = "/api/v2"
	// apiGroup := ctx.R.Group("/api/v2")
	// jwtMiddleware := middleware.NewJwtMiddlewareFactory(ctx.JwtService, nil).GetMiddleware()
	http, _ := ctx.HttpSvc.(iface.DefaultHttpSPI)
	engine := http.GetEngine()

	wgHandler := rest.NewWireguardHandler(ctx.Wireguard)
	wgGroup := engine.Group("/api/wg")
	wgGroup.GET("/clients", wgHandler.GetClients)
}
