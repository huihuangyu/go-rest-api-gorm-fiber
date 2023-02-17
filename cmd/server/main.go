package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/huihuangyu/go-rest-api-gorm-fiber/internal/transport/http"
)

// App - the struct which contains things like the pointers to database connections
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Starting up our application.")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

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
