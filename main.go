package main

import (
	"fmt"
	"mail/src/config"
	"mail/src/routes"
)

func main() {
	appVersion := "v0.5.0"
	fmt.Println("Go Mail Server " + appVersion)

	config := config.Init()
	routes.InitRoutes(config.Port, config.Email, config.Password, appVersion)
}
