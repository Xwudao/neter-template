package main

import (
	"github.com/Xwudao/neter-template/internal/cmd"
)

func main() {
	app, cleanup, err := cmd.MainApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Execute()
	// app.e
	// err = cmd.Execute(func(command *cobra.Command, args []string) {
	// 	err = app.Run()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// })
	// if err != nil {
	// 	panic(err)
	// }
}
