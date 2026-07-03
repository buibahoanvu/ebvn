package api

import (
	"github.com/buibahoanvu/ebvn/internal/handler"
	"github.com/buibahoanvu/ebvn/internal/service"
	"github.com/gin-gonic/gin"
)

type Engine interface {
	 Start() error
}

type engine struct {
	app *gin.Engine
}

func NewEngine() Engine {
	app :=  &engine{
		app: gin.Default(),
	}
	app.initRoutes()

	return app
}

func (e *engine) Start() error{
	return e.app.Run(":8080")
}

func (e *engine) initRoutes() {
	// genpass handler
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)

	e.app.GET("/getpass", genPassHandler.GeneratePassword)
}