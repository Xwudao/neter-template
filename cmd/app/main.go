package main

import (
	"github.com/spf13/cobra"

	"github.com/Xwudao/neter-template/internal/cmd"
)

func main() {

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
