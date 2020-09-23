package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/iamaul/evonix-backend-api/app/models"
	"github.com/iamaul/evonix-backend-api/app/user"

	"github.com/labstack/echo"
)

type response struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type UserHandler struct {
	Usercase user.Usecase
}

func NewUserHandler(e *echo.Echo, uu user.Usecase) {
	handler := &UserHandler{
		Usercase: uu,
	}

	e.POST("/users", handler.Store)
	e.GET("/users/:id", handler.GetByID)
}

func (uh *UserHandler) GetByID(c echo.Context) error {
	paramId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Success: false,
		})
	}

	id := int64(paramId)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := uh.Usercase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &response{
		Code:    http.StatusOK,
		Data:    user,
		Success: true,
	})
}

func createUserValidation(um *models.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(um)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (uh *UserHandler) Store(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &response{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
			Success: false,
		})
	}

	if ok, err := createUserValidation(&user); !ok {
		return c.JSON(http.StatusBadRequest, &response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = uh.Usercase.Store(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusCreated, &response{
		Code:    http.StatusCreated,
		Message: "Successfully created",
		Data:    user,
		Success: true,
	})
}
