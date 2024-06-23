package handler

import (
	"context"
	_interface "myapp/interface"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService _interface.UserRepository
}

func NewUserHandler(userService _interface.UserRepository) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.userService.CreateUser(context.Background(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")

	user, err := h.userService.GetUserById(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
