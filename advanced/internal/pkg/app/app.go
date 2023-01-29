package app

import (
	"advanced/internal/app/endpoint"
	"advanced/internal/app/middleware"
	"advanced/internal/app/service"
	"github.com/labstack/echo/v4"
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
