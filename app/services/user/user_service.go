package users_service

import (
	users_model "go-backend/app/models/users"
	exception_configs "go-backend/internal/configs/exception"
	"go-backend/internal/constants"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	CreateUser(c fiber.Ctx) error
	DeleteUser(c fiber.Ctx) error
	GetAllUser(c fiber.Ctx) error
	GetUserById(c fiber.Ctx) error
}

type UserService struct {
}

func (s *UserService) GetAllUser(c fiber.Ctx) error {
	users := new([]users_model.User)
	jsonResponse := c.Locals(constants.JSONResponse).(func(c fiber.Ctx, status int, message string, data interface{}) error)
	db := c.Locals(constants.Db).(*gorm.DB)

	if err := db.Find(&users).Error; err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Failed to get users", nil)
	}

	var usersWithoutPassword []map[string]interface{}
	for _, user := range *users {
		userMap := map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		}
		usersWithoutPassword = append(usersWithoutPassword, userMap)
	}

	return jsonResponse(c, fiber.StatusOK, "Get all users successfully", usersWithoutPassword)
}

func (s *UserService) GetUserById(c fiber.Ctx) error {
	id := c.Params("id")
	user := new(users_model.User)
	jsonResponse := c.Locals(constants.JSONResponse).(func(c fiber.Ctx, status int, message string, data interface{}) error)
	db := c.Locals(constants.Db).(*gorm.DB)

	if err := db.First(&user, id).Error; err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "User not found", nil)
	}

	return jsonResponse(c, fiber.StatusOK, "Get user by id sucessfully", &user)
}

func (s *UserService) CreateUser(c fiber.Ctx) error {
	user := new(users_model.UserDTO)
	cv := c.Locals(constants.CustomValidator).(exception_configs.CustomValidator)
	jsonResponse := c.Locals(constants.JSONResponse).(func(c fiber.Ctx, status int, message string, data interface{}) error)
	db := c.Locals(constants.Db).(*gorm.DB)

	if err := c.Bind().Body(user); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Params is not empty", nil)
	}

	errors := cv.Validate(user)

	if len(errors) > 0 {
		return jsonResponse(c, fiber.StatusBadRequest, "Params is not empty", errors)
	}

	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Failed to hash password", nil)
	}

	newUser := users_model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return jsonResponse(c, fiber.StatusCreated, "User created sucessfully", &fiber.Map{
		"id":   newUser.ID,
		"name": newUser.Name,
	})
}

func (s *UserService) DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	user := new(users_model.UserDTO)
	jsonResponse := c.Locals(constants.JSONResponse).(func(c fiber.Ctx, status int, message string, data interface{}) error)
	db := c.Locals(constants.Db).(*gorm.DB)

	if err := db.First(&user, id).Error; err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "User not found", nil)
	}

	if err := db.Delete(&user).Error; err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Failed to delete user", nil)
	}

	return jsonResponse(c, fiber.StatusOK, "User deleted sucessfully", nil)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
