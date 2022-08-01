package http

import (
	"github.com/eznd-otus-msa/hw3/app/internal/domain"
	"github.com/eznd-otus-msa/hw3/app/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewPatchUser(pu service.UserPartialUpdater) *patchUserHandler {
	return &patchUserHandler{
		partialUpdater: pu,
	}
}

type patchUserHandler struct {
	partialUpdater service.UserPartialUpdater
}

func (h *patchUserHandler) Handle() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := ctx.ParamsInt(UserIdFieldName, 0)
		if err != nil {
			return fail(ctx, err)
		}

		var pu service.UserPartialUpdate
		err = ctx.BodyParser(&pu)
		if err != nil {
			return fail(ctx, err)
		}

		user, err := h.partialUpdater.PartialUpdate(domain.UserId(userId), pu.ToDomain())
		if err != nil {
			return fail(ctx, err)
		}
		return json(ctx, (&service.User{}).FromDomain(user))
	}
}
