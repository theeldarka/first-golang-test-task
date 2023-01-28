package main

import (
	"github.com/labstack/gommon/log"
	"simple/internal/pkg/app"
)

func main() {
	a := app.New()

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
