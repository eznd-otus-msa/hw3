package http

import (
	"github.com/gofiber/fiber/v2"
	"math/rand"
)

const FailureChance = 10

func NewChaosMonkeyMw() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if rand.Intn(99) < FailureChance {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.Next()
	}
}
