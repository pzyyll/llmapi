package service

import (
	"fmt"
	"regexp"
	"time"

	"llmapi/src/internal/config"
	"llmapi/src/internal/constants"
	"llmapi/src/internal/model"
	"llmapi/src/internal/repository"
	"llmapi/src/internal/utils"
	"llmapi/src/pkg/auth"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(username, password string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	GetUserByName(username string) (*model.User, error)
	DeleteUser(id int64) error
	InitAdminUser() error
}

type userService struct {
	userRepo     repository.UserRepo
	uidGenerator utils.UidGenerator
	cfg          *config.Config
}

func NewUserService(userRepo repository.UserRepo, cfg *config.Config, uidGenerator utils.UidGenerator) UserService {
	return &userService{
		userRepo:     userRepo,
		cfg:          cfg,
		uidGenerator: uidGenerator,
	}
}

func (s *userService) CreateUser(username, password string) (*model.User, error) {
	if _, err := s.userRepo.GetUserByName(username); err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	if len(password) < constants.PasswordMinLength || len(password) > constants.PasswordMaxLength {
		return nil, fmt.Errorf("password length must be between %d and %d characters", constants.PasswordMinLength, constants.PasswordMaxLength)
	}

	if len(username) < constants.UsernameMinLength || len(username) > constants.UsernameMaxLength {
		return nil, fmt.Errorf("username length must be between %d and %d characters", constants.UsernameMinLength, constants.UsernameMaxLength)
	}

	if !regexp.MustCompile(constants.UsernameRegex).MatchString(username) {
		return nil, fmt.Errorf("username can only contain alphanumeric characters and underscores")
	}

	user := &model.User{
		UserID:   s.uidGenerator.GenerateUID(),
		Username: username,
		Password: password,
		IsActive: true,
		Role:     constants.RoleTypeUser,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(id int64) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) GetUserByName(username string) (*model.User, error) {
	return s.userRepo.GetUserByName(username)
}

func (s *userService) DeleteUser(id int64) error {
	return s.userRepo.DeleteUser(id)
}

func (s *userService) InitAdminUser() error {
	_, err := s.userRepo.FindFirstUserByRole(constants.RoleTypeSuper)
	if err == nil {
		return fmt.Errorf("Admin user is init")
	}

	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check if admin user exists: %w", err)
	}

	hashedPassword, err := auth.HashPassword(s.cfg.AdminPassword)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}

	adminUser := &model.User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ID:        uint(s.uidGenerator.GenerateUID()),
		},
		Username: s.cfg.AdminUser,
		Password: hashedPassword,
		IsActive: true,
		Role:     constants.RoleTypeSuper,
	}

	if err := s.userRepo.CreateUser(adminUser); err != nil {
		return err
	}

	return nil
}
