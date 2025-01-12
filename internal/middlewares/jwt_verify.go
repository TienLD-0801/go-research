package middlewares

import (
	"fmt"
	"go-backend/internal/constants"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64 `json:"id"`
	jwt.RegisteredClaims
}

func VerifyToken(c fiber.Ctx) error {
	jsonResponse := c.Locals(constants.JSONResponse).(func(c fiber.Ctx, status int, message string, data interface{}) error)

	tokenString := c.Get("Authorization")
	fmt.Println("Authorization header:", tokenString)

	if tokenString == "" {
		return jsonResponse(c, fiber.StatusUnauthorized, "Token is required", nil)
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return jsonResponse(c, fiber.StatusUnauthorized, "Invalid token", nil)
	}

	if !token.Valid {
		fmt.Println("Invalid token:", token)
		return jsonResponse(c, fiber.StatusUnauthorized, "Invalid token", nil)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		fmt.Println("Invalid token claims:", token.Claims)
		return jsonResponse(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
	}

	c.Locals(constants.VerifyToken, claims.UserID)

	return c.Next()
}
