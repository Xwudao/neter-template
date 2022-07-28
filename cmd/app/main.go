package main

import (
	"github.com/Xwudao/neter-template/internal/cmd"
	"github.com/spf13/cobra"
)

func main() {

	err := cmd.Execute(func(cmd *cobra.Command, args []string) {
		app, cleanup, err := mainApp()
		if err != nil {
			panic(err)
		}

		defer cleanup()

		err = app.Run()
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
}
