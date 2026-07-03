package main

import "github.com/buibahoanvu/ebvn/internal/api"

func main() {
	app := api.NewEngine()
	err := app.Start()
	if err != nil {
		panic(err)
	}
}
