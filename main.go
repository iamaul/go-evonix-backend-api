package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iamaul/evonix-backend-api/app/middleware"
	udh "github.com/iamaul/evonix-backend-api/app/user/delivery/http"
	ur "github.com/iamaul/evonix-backend-api/app/user/repository"
	uu "github.com/iamaul/evonix-backend-api/app/user/usecase"
	"github.com/iamaul/evonix-backend-api/config"
	"github.com/iamaul/evonix-backend-api/config/database"

	"github.com/labstack/echo"
)

func main() {
	config := config.NewConfig()

	connection, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	appMiddl := middleware.InitAppMiddleware(config.AppName)

	e.Use(appMiddl.CORS)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusAccepted,
			"message": "EvoniX Backend API v1.0",
		})
	})

	userRepo := ur.NewUserRepository(connection.SQL)

	timeoutContext := time.Duration(2) * time.Second

	userCase := uu.NewUserUsecase(userRepo, timeoutContext)

	udh.NewUserHandler(e, userCase)

	log.Fatal(e.Start(fmt.Sprintf(`%s`, config.AppPort)))
}
