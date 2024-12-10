package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(ctx *fiber.Ctx) error {
	start := time.Now()
	err := ctx.Next()
	log.Printf("[%s] %s %s %d %s",
		time.Now().Format(time.RFC3339),
		ctx.Method(),
		ctx.Path(),
		ctx.Response().StatusCode(),
		time.Since(start),
	)
	return err
}
