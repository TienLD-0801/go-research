package auth_service

import "github.com/gofiber/fiber/v3"

type IAuthService interface {
	Login(c fiber.Ctx) error
}

type AuthService struct {
}

func (s *AuthService) Login(c fiber.Ctx) error {
	return nil
}
