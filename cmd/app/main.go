package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/Xwudao/neter-template/internal/cmd"
)

var (
	buildTime = ""
	// shutdownTimeout bounds the release of construction-time resources.
	shutdownTimeout = 10 * time.Second
)

func main() {
	fmt.Println("app build time: ", buildTime)
	err := cmd.Execute(func(cmd *cobra.Command, args []string) {
		app, lifecycle, err := mainApp()
		if err != nil {
			panic(err)
		}
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()
			if err := lifecycle.Stop(ctx); err != nil {
				fmt.Println("cleanup:", err)
			}
		}()

		if err := app.Prepare(); err != nil {
			app.AbortStartup()
			panic(err)
		}
		if err := lifecycle.Start(context.Background()); err != nil {
			app.CancelAppContext()
			panic(err)
		}

		runCtx, stopSignals := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stopSignals()
		if err := app.Wait(runCtx); err != nil {
			panic(err)
		}

	})
	if err != nil {
		panic(err)
	}
}
