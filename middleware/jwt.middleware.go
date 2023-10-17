package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Jwt(ctx *fiber.Ctx) error {
	fmt.Println("ini middleware")
	return ctx.Next()
}
