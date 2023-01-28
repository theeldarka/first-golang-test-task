package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strings"
)

func CheckUserRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		role := ctx.Request().Header.Get("User-Role")

		if strings.Contains(strings.ToLower(role), "admin") {
			log.Warn("red button user detected")
		}

		return next(ctx)
	}
}
