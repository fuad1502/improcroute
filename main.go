package main

import (
	"github.com/fuad1502/improcroute/service"
)

func main() {
	var service service.ImprocrouteService
	service.Start(":8080")
}
