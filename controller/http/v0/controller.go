package httpv0

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type userUC interface {
	UpdateUsername(ctx context.Context, id int64, newName string) error
}

type userHandler struct {
	uc userUC
}

func NewUserHandler(uc userUC) *userHandler {
	return &userHandler{uc: uc}
}

func (h *userHandler) UpdateUsername(c *fiber.Ctx) error {
	id := int64(1)
	name := "new_username"

	err := h.uc.UpdateUsername(c.Context(), id, name)
	if err != nil {
		log.Printf("failed update username: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
