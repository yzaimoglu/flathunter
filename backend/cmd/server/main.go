package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/http/api"
)

func main() {
	if !fiber.IsChild() {
		config.Load()
		config.SetupLogger()
		config.SetupArango()
	}

	server := fiber.New(
		fiber.Config{
			Prefork: true,
		},
	)
	server.Use(helmet.New())

	file, err := os.OpenFile("./logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Fatal("Error opening file: %v", err)
	}
	defer file.Close()

	server.Use(logger.New(logger.Config{
		Format:     "[${time}] [http] [${status}] ${method} ${path} - ${ip} | ${latency}\n",
		TimeFormat: "02.01.2006T15:04:05",
		TimeZone:   "Europe/Berlin",
		Output:     file,
	}))
	server.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
		URL:  "/favicon.ico",
	}))

	api.APIv1(server)

	go func() {
		if err := server.Listen(":" + config.GetString("SERVER_PORT")); err != nil {
			slog.Fatalf("Error starting server: %v", err)
		}
	}()

	config.SysCallSetup()
}
