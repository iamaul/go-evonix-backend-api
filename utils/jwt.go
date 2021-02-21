package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iamaul/evonix-backend-api/app/models"
)

// AccessTokenClaims will claims data user by validating their token
type AccessTokenClaims struct {
	UserID int64 `json:"uid"`
	Admin  uint8 `json:"admin"`
	Helper uint8 `json:"helper"`
	jwt.StandardClaims
}

// RefreshTokenClaims allows user to obtain a new JWT
type RefreshTokenClaims struct {
	UserID int64 `json:"uid"`
	jwt.StandardClaims
}

// TokenPayload will generate data user of token payload
type TokenPayload struct {
	UserID int64
	Admin  uint8
	Helper uint8
}

// GenerateAccessToken will generate access JWT token
func GenerateAccessToken(user models.User) string {
	claims := AccessTokenClaims{
		user.ID,
		user.Admin,
		user.Helper,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // 15 minutes
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	res, _ := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

	return res
}

// VerifyAccessToken verify access JWT token
func VerifyAccessToken(accessToken string) (interface{}, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(InvalidSigningMethodErr)
		}

		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payload := TokenPayload{
			UserID: claims["uid"].(int64),
			Admin:  claims["admin"].(uint8),
			Helper: claims["helper"].(uint8),
		}
		return payload, nil
	}
	return nil, errors.New(TokenErr)
}

// GenerateRefreshToken will generate a new JWT refresh token
func GenerateRefreshToken(user models.User) string {
	claims := RefreshTokenClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	res, _ := token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))

	return res
}

// VerifyRefreshToken verify new JWT refresh token
func VerifyRefreshToken(refreshToken string) (interface{}, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(InvalidSigningMethodErr)
		}

		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["uid"].(int64)
		return userID, nil
	}

	return nil, errors.New(InvalidRefreshTokenErr)
}
