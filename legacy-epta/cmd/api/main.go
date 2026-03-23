package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "backend/docs"
	"backend/internal/app"
)

// @title Student Schedule API
// @version 1.0
// @description API для системы расписания студентов
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http

func main() {
	application := app.New()

	if err := application.Init(); err != nil {
		log.Fatal("Failed to initialize application:", err)
	}

	// канал для обработки сигналов завершения
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := application.Run(); err != nil {
			log.Fatal("Failed to run application:", err)
		}
	}()

	log.Println("Server started successfully")

	<-done
	log.Println("Shutting down server...")

	application.Close()
	log.Println("Server stopped")
}
