package main

import (
	"conductor/api/controller"
	"conductor/db"
	"conductor/job"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	db.Init()
	controller.Init(app)
	job.Init(context.Background())

	defer db.Close()

	log.Fatal(app.Listen("0.0.0.0:9091"))
}
