package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysFromNow() int
}

type Endpoint struct {
	svc Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{svc: svc}
}

func (e *Endpoint) DaysFromNow(ctx echo.Context) error {
	response := fmt.Sprintf("There are %d days until 2025", e.svc.DaysFromNow())

	return ctx.String(http.StatusOK, response)
}
