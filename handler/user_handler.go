package handler

import (
	"strconv"

	"github.com/Mr-Chegini/simple-golang-crud/models"
	"github.com/Mr-Chegini/simple-golang-crud/service"
	"github.com/gofiber/fiber/v2"
)

type userCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type errorResponse struct {
	Message string `json:"message"`
}

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", h.ListUsers)
	users.Post("/", h.CreateUser)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var payload userCreateRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid request body"})
	}

	user := models.User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	if err := h.userService.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userService.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Message: err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid user id"})
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Message: err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid user id"})
	}

	var payload userUpdateRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid request body"})
	}

	user := models.User{
		ID:    uint(id),
		Name:  payload.Name,
		Email: payload.Email,
	}

	if err := h.userService.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: "invalid user id"})
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Message: err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
