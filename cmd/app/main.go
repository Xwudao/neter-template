package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Xwudao/neter-template/internal/cmd"
)

var (
	buildTime = ""
)

func main() {
	fmt.Println("app build time: ", buildTime)
	err := cmd.Execute(func(cmd *cobra.Command, args []string) {
		app, cleanup, err := mainApp()
		if err != nil {
			panic(err)
		}
		err = app.Run()
		if err != nil {
			panic(err)
		}

		defer cleanup()

	})
	if err != nil {
		panic(err)
	}
}
