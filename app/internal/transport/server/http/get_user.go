package http

import (
	"github.com/eznd-otus-msa/hw3/app/internal/domain"
	"github.com/eznd-otus-msa/hw3/app/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewGetUser(r service.UserReader) *getUserHandler {
	return &getUserHandler{
		reader: r,
	}
}

type getUserHandler struct {
	reader service.UserReader
}

func (h *getUserHandler) Handle() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := ctx.ParamsInt(UserIdFieldName, 0)
		if err != nil {
			return fail(ctx, err)
		}

		u, err := h.reader.Get(domain.UserId(userId))
		if err != nil {
			return fail(ctx, err)
		}
		return json(ctx, (&service.User{}).FromDomain(u))
	}
}
