package logic

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"

	"wmjtyd-iot/internal/module/auth/model"
	"wmjtyd-iot/internal/module/auth/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize      = 16
	iterations    = 10000
	keyLength     = 32
	tokenDuration = 24 * time.Hour
	secretKey     = "your-secret-key"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type AuthLogic struct {
	userRepo *repository.UserRepository
}

func NewAuthLogic(userRepo *repository.UserRepository) *AuthLogic {
	return &AuthLogic{userRepo: userRepo}
}

func generateSalt() (string, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func hashPassword(password, salt string) string {
	saltBytes, _ := base64.StdEncoding.DecodeString(salt)
	hash := pbkdf2.Key([]byte(password), saltBytes, iterations, keyLength, sha256.New)
	return base64.StdEncoding.EncodeToString(hash)
}

func verifyPassword(password, hash, salt string) bool {
	return hashPassword(password, salt) == hash
}

func generateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(tokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func (l *AuthLogic) HandleLogin(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid request"})
	}

	user, err := l.userRepo.GetByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Internal server error"})
	}
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{Message: "Invalid credentials"})
	}

	if !verifyPassword(req.Password, user.Password, user.Salt) {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{Message: "Invalid credentials"})
	}

	token, err := generateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to generate token"})
	}

	return c.JSON(LoginResponse{Token: token})
}

func (l *AuthLogic) HandleCreateUser(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid request"})
	}

	existingUser, err := l.userRepo.GetByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Internal server error"})
	}
	if existingUser != nil {
		return c.Status(fiber.StatusConflict).JSON(ErrorResponse{Message: "Username already exists"})
	}

	salt, err := generateSalt()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to generate salt"})
	}
	passwordHash := hashPassword(req.Password, salt)

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
		Salt:     salt,
		Status:   "1",
	}

	if err := l.userRepo.Create(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to create user"})
	}

	return c.JSON(RegisterResponse{ID: string(user.ID)})
}

func (l *AuthLogic) HandleRegister(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid request"})
	}

	existingUser, err := l.userRepo.GetByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Internal server error"})
	}
	if existingUser != nil {
		return c.Status(fiber.StatusConflict).JSON(ErrorResponse{Message: "Username already exists"})
	}

	salt, err := generateSalt()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to generate salt"})
	}
	passwordHash := hashPassword(req.Password, salt)

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
		Salt:     salt,
		Status:   "1",
	}

	if err := l.userRepo.Create(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to create user"})
	}

	return c.JSON(RegisterResponse{ID: string(user.ID)})
}

func (l *AuthLogic) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "User ID is required"})
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid user ID"})
	}

	user, err := l.userRepo.GetByID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Internal server error"})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Message: "User not found"})
	}

	return c.JSON(user)
}

func (l *AuthLogic) HandleUpdateUser(c *fiber.Ctx) error {
	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid request"})
	}

	userID, err := strconv.Atoi(req.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid user ID"})
	}

	user, err := l.userRepo.GetByID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Internal server error"})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Message: "User not found"})
	}

	user.Username = req.Username
	user.Email = req.Email

	if err := l.userRepo.Update(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to update user"})
	}

	return c.JSON(user)
}

func (l *AuthLogic) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "User ID is required"})
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid user ID"})
	}

	if err := l.userRepo.Delete(uint(userID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to delete user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (l *AuthLogic) HandleDisableUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "User ID is required"})
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid user ID"})
	}

	user, err := l.userRepo.GetByID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Internal server error"})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Message: "User not found"})
	}

	user.Status = "0"
	if err := l.userRepo.Update(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to disable user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (l *AuthLogic) HandleEnableUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "User ID is required"})
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: "Invalid user ID"})
	}

	user, err := l.userRepo.GetByID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Internal server error"})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Message: "User not found"})
	}

	user.Status = "1"
	if err := l.userRepo.Update(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to enable user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (l *AuthLogic) HandleGetUsers(c *fiber.Ctx) error {
	users, err := l.userRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: "Failed to get users"})
	}

	return c.JSON(users)
}
