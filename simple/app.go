package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
	"time"
)

func daysFromNow() int {
	year := 2025
	month := time.January
	day := 1

	now := time.Now()
	then := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	days := then.Sub(now).Hours() / 24

	return int(days)
}

func main() {
	e := echo.New()

	e.GET("/", daysHandler, checkUserRoleMiddleware)

	e.Logger.Fatal(e.Start(":1323"))

	fmt.Println(daysFromNow())
}

func daysHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("There are %d days until 2025", daysFromNow()))
}

func checkUserRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Header.Get("User-Role")

		if strings.Contains(role, "admin") {
			log.Warn("red button user detected")
		}

		return next(c)
	}
}
