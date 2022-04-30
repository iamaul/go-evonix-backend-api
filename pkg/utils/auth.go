package utils

import (
	"fmt"
	"strings"

	"github.com/iamaul/go-evonix-backend-api/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetCurrentUserUUIDFROMJWT(c echo.Context) uuid.UUID {
	bearerHeader := c.Request().Header.Get("Authorization")
	UUIDZero := uuid.NullUUID{
		UUID:  uuid.UUID{},
		Valid: false,
	}
	if bearerHeader != "" {
		headerParts := strings.Split(bearerHeader, " ")
		if len(headerParts) != 2 {
			return UUIDZero.UUID
		}

		tokenString := headerParts[1]
		if tokenString == "" {
			return UUIDZero.UUID
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
			}
			cfgGlobal := config.GetConfig()
			secret := []byte(cfgGlobal.Server.JwtSecretKey)
			return secret, nil
		})
		if err != nil {
			return UUIDZero.UUID
		}

		if !token.Valid {
			return UUIDZero.UUID
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		customerId := claims["sub"].(string)
		if ok && token.Valid && customerId != "" {
			customerUUID, err := uuid.Parse(customerId)
			if err != nil {
				return UUIDZero.UUID
			}
			return customerUUID

		} else if !ok || customerId == "" {
			return UUIDZero.UUID
		}
	}
	return UUIDZero.UUID
}
