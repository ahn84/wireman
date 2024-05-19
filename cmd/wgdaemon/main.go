package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"wgd/src/app"
	"wgd/src/configuration"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "go-app",
	Short: "Default cmd short description",
	Long:  `Default cmd long description`,
	Run:   rootCmdRun,
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	runCtx := context.Background()
	app.InitAppContext()
	app.AppStart(runCtx)

	// handle signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		log.Println("Shutting down server...")
		if err := app.AppShutdown(ctx); err != nil {
			log.Fatalf("Forcefully shutdown: %v\n", err)
		}
		done <- true
	}()

	<-done
}

func init() {
	cobra.OnInitialize(configuration.InitializeAppConfig)
}
