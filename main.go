package main

import (
	"fmt"
	"mail/src/config"
	"mail/src/routes"
)

func main() {
	appVersion := "v0.4.0"
	fmt.Println("Go Mail Server " + appVersion)

	config := config.Init()
	fmt.Println(config)
	routes.InitRoutes(config.Port, config.Email, config.Password, appVersion)
}
