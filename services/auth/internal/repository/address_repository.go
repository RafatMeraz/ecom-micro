package repository

import (
	"github.com/RafatMeraz/ecom-micro/auth/internal/model"
	ce "github.com/RafatMeraz/ecom-micro/pkg/errors"
	"gorm.io/gorm"
	"log/slog"
)

type AddressRepository interface {
	CreateAddress(model.Address) (model.Address, error)
	GetAddressByUserId(userId uint) ([]model.Address, error)
}

type AddressRepositoryImpl struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{
		db: db,
	}
}

func (a AddressRepositoryImpl) CreateAddress(address model.Address) (model.Address, error) {
	result := a.db.Create(&address)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return address, ce.ErrDatabaseOperation
	}
	return address, nil
}

func (a AddressRepositoryImpl) GetAddressByUserId(userId uint) ([]model.Address, error) {
	var addresses []model.Address
	err := a.db.Where("user_id = ?", userId).Find(&addresses).Error
	if err != nil {
		slog.Error(err.Error())
		return addresses, ce.ErrDatabaseOperation
	}
	return addresses, nil
}
