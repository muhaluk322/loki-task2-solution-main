package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"path/filepath"
	"webapp/pkg/config"
)

func main() {
	logfilePath := config.GetEnv("LOG_PATH", "app.log")
	logDir := filepath.Dir(logfilePath)
	err := os.MkdirAll(logDir, os.ModePerm)

	if err != nil {
		fmt.Println("Failed to create directories:", err)
		return
	}

	logfile, _ := os.Create(logfilePath)
	port := config.GetEnv("PORT", "3000")

	e := echo.New()
	e.Static("/", "public")
	e.Use(middleware.Logger())
	e.Logger.SetOutput(logfile)

	e.GET("/", func(c echo.Context) error {
		return c.File("public/views/webapp.html")
	})
	e.Start(":" + port)
}
