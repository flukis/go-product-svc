package handler

import (
	"net/http"
	"time"

	"github.com/fahmilukis/go-product-svc/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	ProductUC domain.ProductUsecase
}

func ProductRoute(a *fiber.App, puc domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUC: puc,
	}

	route := a.Group("/api/v1")

	route.Post("/product", handler.CreateProduct)
}

func (puc *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	// check kelengkapan formData
	prd := &domain.Products{}
	if err := c.BodyParser(prd); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"msg":    err.Error(),
		})
	}
	// insert product ke db
	now := time.Now()
	id := uuid.NewString()

	prd.ID = id
	prd.CreatedAt = now
	prd.UpdatedAt = now

	if err := puc.ProductUC.Store(c.Context(), prd); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"msg":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "success add product",
	})
}
