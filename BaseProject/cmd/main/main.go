package main

import (
	"Examples/BaseProject/internal/config"
	"Examples/BaseProject/internal/handlers"
	"Examples/BaseProject/internal/logger"
	"Examples/BaseProject/internal/mail"
	"Examples/BaseProject/internal/models"
	"Examples/BaseProject/internal/routes"
	"Examples/BaseProject/pkg/env"

	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfig(env.GetValue("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	if err := logger.SetupLogging(cfg.LogFile); err != nil {
		log.Fatal(err)
	}

	models.New(cfg)
	logger.Info("База данных подключена")

	mail.New(cfg)
	logger.Info("Почта подключена")

	e := echo.New()

	e.Use(middleware.Recover())

	e.Static("/static", filepath.Join("static"))

	renderer, err := handlers.NewTemplate(filepath.Join("internal", "templates", "*.html"))
	if err != nil {
		logger.Warn("Ошибка загрузки шаблонов: ", err)
	}
	e.Renderer = renderer

	routes.SetupRoutes(e)

	logger.Infof("Сервер запущен на порту %s", cfg.ServerPort)
	if err := e.Start(cfg.ServerPort); err != nil && err != http.ErrServerClosed {
		logger.Error("Ошибка запуска сервера: ", err)
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}