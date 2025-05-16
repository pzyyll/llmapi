package repository

import (
	"llmapi/src/internal/constants"
	"llmapi/src/internal/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	// CreateUser creates a new user in the database
	CreateUser(user *model.User) error
	// GetUserByID retrieves a user by their ID
	GetUserByID(id int64) (*model.User, error)
	GetUserByUserID(userID int64) (*model.User, error)
	GetUserByName(username string) (*model.User, error)
	GetUsers() (*[]model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
	FindFirstUserByRole(role constants.RoleType) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByUserID(userID int64) (*model.User, error) {
	var user model.User
	if err := r.db.Where(&model.User{UserID: userID}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByName(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where(&model.User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) DeleteUser(user *model.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) FindFirstUserByRole(role constants.RoleType) (*model.User, error) {
	var user model.User
	if err := r.db.Where(&model.User{Role: string(role)}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUsers() (*[]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
