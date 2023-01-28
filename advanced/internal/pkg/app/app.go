package app

import (
	"github.com/labstack/echo/v4"
	"simple/internal/app/endpoint"
	"simple/internal/app/middleware"
	"simple/internal/app/service"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	server   *echo.Echo
}

func New() *App {
	app := &App{}

	app.service = service.New()
	app.endpoint = endpoint.New(app.service)
	app.server = echo.New()

	app.server.GET("/", app.endpoint.DaysFromNow, middleware.CheckUserRole)

	return app
}

func (a *App) Run() error {
	return a.server.Start(":1323")
}
