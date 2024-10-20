package repository

import (
	"errors"
	"github.com/RafatMeraz/ecom-micro/auth/internal/model"
	ce "github.com/RafatMeraz/ecom-micro/pkg/errors"
	"gorm.io/gorm"
	"log/slog"
)

type UserRepository interface {
	CreateUser(model.User) (model.User, error)
	DeleteUser(uint) error
	GetUserByEmail(string) (model.User, error)
	GetUserByID(uint) (model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) CreateUser(user model.User) (model.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return user, result.Error
	}
	return user, nil
}

func (u UserRepositoryImpl) DeleteUser(userId uint) error {
	result := u.db.Delete(&model.User{}, userId)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ce.ErrRecordNotFound
		}
		return ce.ErrDatabaseOperation
	}
	return nil
}

func (u UserRepositoryImpl) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, ce.ErrRecordNotFound
		}
		return user, ce.ErrDatabaseOperation
	}
	return user, nil
}

func (u UserRepositoryImpl) GetUserByID(userId uint) (model.User, error) {
	var user model.User
	result := u.db.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, ce.ErrRecordNotFound
		}
		return user, ce.ErrDatabaseOperation
	}
	return user, nil
}
