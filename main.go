package main

import (
	"fmt"
	"mail/src/config"
	"mail/src/routes"
)

func main() {
	appVersion := "v0.5.0"
	fmt.Println("Go Mail Server " + appVersion)

	appConfig := config.Init()
	routes.InitRoutes(appConfig.Port, appConfig.Email, appConfig.Password, appVersion)
}
