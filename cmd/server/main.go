package main

import "fmt"

// App - the struct which contains things like the pointers to database connections
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Starting up our application.")
	return nil
}

func main() {
	fmt.Println("GO REEST API Practise")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
