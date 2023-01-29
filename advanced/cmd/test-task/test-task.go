package main

import (
	"advanced/internal/pkg/app"
	"github.com/labstack/gommon/log"
)

func main() {
	a := app.New()

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
