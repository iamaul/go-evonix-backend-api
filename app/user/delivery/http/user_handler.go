package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/iamaul/evonix-backend-api/app/models"
	"github.com/iamaul/evonix-backend-api/app/user"
	"github.com/iamaul/evonix-backend-api/utils"

	"github.com/labstack/echo"
)

type UserHandler struct {
	Usercase user.Usecase
}

func NewUserHandler(e *echo.Echo, uu user.Usecase) {
	handler := &UserHandler{
		Usercase: uu,
	}

	g := e.Group("/api/v1")
	g.POST("/users", handler.Store)
	g.GET("/users/:id", handler.GetByID)
}

func (uh *UserHandler) GetByID(c echo.Context) error {
	uID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &utils.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Success: false,
		})
	}

	id := int64(uID)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := uh.Usercase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &utils.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.Response{
		Code:    http.StatusOK,
		Result:  user,
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
		return c.JSON(http.StatusUnprocessableEntity, &utils.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
			Success: false,
		})
	}

	if ok, err := createUserValidation(&user); !ok {
		return c.JSON(http.StatusBadRequest, &utils.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user.RegisteredDate = time.Now().Unix()
	user.RegisterIP = c.RealIP()

	err = uh.Usercase.Store(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})
	}

	token := utils.GenerateAccessToken(user)

	return c.JSON(http.StatusCreated, &utils.Response{
		Code:        http.StatusCreated,
		Message:     "User created successfully",
		Result:      user,
		AccessToken: token,
		Success:     true,
	})
}
