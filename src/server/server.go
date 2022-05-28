package server

import (
	"fmt"
	"os"

	"github.com/ELPanaJose/api-deno-compiler/src/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	HerokuEchoIpDashboard "github.com/paij0se/heroku-echo-ip-dashboard/src"
)

func StartServer() {
	e := echo.New()
	HerokuEchoIpDashboard.HerokuEchoIpDashboard(e) // init the dashboard
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/", "src/public")
	e.GET("/code", routes.GetCode)
	e.POST("/code", routes.PostCode)
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}
	fmt.Printf("Api on port: %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
