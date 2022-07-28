package main

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	err = app.Execute()
	if err != nil {
		panic(err)
	}
}
