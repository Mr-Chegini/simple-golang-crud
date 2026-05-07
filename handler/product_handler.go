package handler

import (
	"strconv"

	"github.com/Mr-Chegini/simple-golang-crud/models"
	"github.com/Mr-Chegini/simple-golang-crud/service"
	"github.com/gofiber/fiber/v2"
)

type productCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	UserID      uint    `json:"user_id"`
}

type productUpdateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	UserID      uint    `json:"user_id"`
}

type ProductHandler struct {
	productService service.ProductServiceInterface
}

func NewProductHandler(productService service.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) RegisterRoutes(router fiber.Router) {
	products := router.Group("/products")
	products.Get("/", h.ListProducts)
	products.Post("/", h.CreateProduct)
	products.Get("/:id", h.GetProduct)
	products.Put("/:id", h.UpdateProduct)
	products.Delete("/:id", h.DeleteProduct)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var payload productCreateRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid request body"})
	}

	product := models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		UserID:      payload.UserID,
	}

	if err := h.productService.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) ListProducts(c *fiber.Ctx) error {
	products, err := h.productService.ListProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Message: err.Error()})
	}
	return c.JSON(products)
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid product id"})
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Message: err.Error()})
	}

	return c.JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid product id"})
	}

	var payload productUpdateRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid request body"})
	}

	product := models.Product{
		ID:          uint(id),
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		UserID:      payload.UserID,
	}

	if err := h.productService.UpdateProduct(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: err.Error()})
	}

	return c.JSON(product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid product id"})
	}

	if err := h.productService.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Message: err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
