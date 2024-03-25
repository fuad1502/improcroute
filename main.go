package main

import (
	"os"

	"github.com/fuad1502/improcroute/service"
)

func main() {
	var service service.ImprocrouteService
	service.Start(":" + os.Getenv("IPR_PORT"))
}
