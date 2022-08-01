package http

import (
	"github.com/eznd-otus-msa/hw3/app/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewPostUser(c service.UserCreator) *postUserHandler {
	return &postUserHandler{
		creator: c,
	}
}

type postUserHandler struct {
	creator service.UserCreator
}

func (h *postUserHandler) Handle() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var u service.UserCreate
		err := ctx.BodyParser(&u)
		if err != nil {
			return fail(ctx, err)
		}

		user, err := h.creator.Create(&u)
		if err != nil {
			return fail(ctx, err)
		}
		return created(ctx, (&service.User{}).FromDomain(user), int64(user.Id))
	}
}
