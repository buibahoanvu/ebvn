package api

import (
	"fmt"
	"net/http"

	"github.com/buibahoanvu/ebvn/internal/handler"
	"github.com/buibahoanvu/ebvn/internal/service"
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Start() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type engine struct {
	app *gin.Engine
	cfg *Config
}

func NewEngine(cfg *Config) Engine {
	app := &engine{
		app: gin.Default(),
		cfg: cfg,
	}
	app.initRoutes()

	return app
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s",e.cfg.AppPort))
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

func (e *engine) initRoutes() {
	// genpass handler
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)

	e.app.GET("/genpass", genPassHandler.GeneratePassword)
}
