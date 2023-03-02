package handler

import (
	"net/http"
	"time"

	"github.com/fahmilukis/go-product-svc/domain"
	pkg "github.com/fahmilukis/go-product-svc/pkg/utils"
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
	route.Get("/product", handler.GetListProducts)
	route.Get("/product/:id", handler.GetProductDetail)
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

func (puc *ProductHandler) GetListProducts(c *fiber.Ctx) error {
	params := &pkg.Pagination{}
	if err := c.BodyParser(params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"msg":    err.Error(),
		})
	}

	data, nextPagination, err := puc.ProductUC.Fetch(c.Context(), *params)
	if err != nil {
		if err == domain.ErrNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status": false,
				"msg":    err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"msg":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "success get product lists",
		"data":   data,
		"meta":   nextPagination,
	})
}

func (puc *ProductHandler) GetProductDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := puc.ProductUC.GetByID(c.Context(), id)

	if err != nil {
		if err == domain.ErrNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status": false,
				"msg":    err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"msg":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "success get product",
		"data":   data,
	})
}
