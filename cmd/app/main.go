package main

import (
	"context"
	"fmt"
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

		err = app.Run()
		if err != nil {
			panic(err)
		}

	})
	if err != nil {
		panic(err)
	}
}
