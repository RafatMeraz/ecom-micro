package service

import (
	"github.com/RafatMeraz/ecom-micro/auth/internal/dto"
	"github.com/RafatMeraz/ecom-micro/auth/internal/repository"
)

type UserService struct {
	userRepository    repository.UserRepository
	addressRepository repository.AddressRepository
}

func NewUserService(
	userRepository repository.UserRepository,
	addressRepository repository.AddressRepository) *UserService {
	return &UserService{
		userRepository:    userRepository,
		addressRepository: addressRepository,
	}
}

func (service *UserService) SignUp(signUpRequest dto.SignUpRequest) (dto.SignUpResponse, error) {
	user, err := service.userRepository.CreateUser(signUpRequest.ToModel())
	if err != nil {
		return dto.SignUpResponse{}, err
	}

	return dto.NewSignUpResponseFromEntity(user), nil
}
