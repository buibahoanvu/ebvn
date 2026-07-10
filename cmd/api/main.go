package main

import "github.com/buibahoanvu/ebvn/internal/api"

//@title Bookmark Management API
//@version 1.0.0
//@description API for managing bookmarks
//@BasePath /
func main() {
	//Create app config
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}

	app := api.NewEngine(cfg)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
