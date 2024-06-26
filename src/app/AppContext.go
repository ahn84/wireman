package app

import (
	"context"
	"log/slog"
	"wgd/src/configuration"
	"wgd/src/service/impl"

	"sync"
)

var (
	ctx     *configuration.AppContext
	ctxOnce sync.Once
)

func InitAppContext() *configuration.AppContext {
	ctxOnce.Do(func() {
		slog.Info("Initializing ...")
		ctx = &configuration.AppContext{}
		doInit_(ctx)
	})
	return ctx
}

func AppStart(c context.Context) error {
	if ctx == nil || ctx.HttpSvc == nil {
		return nil
	}
	var err error
	if ctx.HttpSvc != nil {
		err = ctx.HttpSvc.Start(c)
	}
	return err
}

func AppShutdown(c context.Context) error {
	if ctx == nil || ctx.HttpSvc == nil {
		return nil
	}
	var err error
	if ctx.HttpSvc != nil {
		err = ctx.HttpSvc.Stop(c)
	}
	return err
}

// doInit_ initialize app components like services and controllers, resolving their dependencies properly
func doInit_(c *configuration.AppContext) {
	ctx := context.Background()
	svc, err := impl.NewWGManager("wg0")
	if err != nil {
		slog.Warn("Init wireguard manager", "error", err)
	} else {
		c.Wireguard = svc
		c.Wireguard.Initialize(ctx)
	}

	c.HttpSvc = impl.NewHttpService(c.Wireguard)
	c.HttpSvc.Initialize(ctx)
}
