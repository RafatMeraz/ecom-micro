package service

import (
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/RafatMeraz/ecom-micro/auth/internal/dto"
	"github.com/RafatMeraz/ecom-micro/auth/internal/repository"
	ce "github.com/RafatMeraz/ecom-micro/pkg/errors"
	"github.com/RafatMeraz/ecom-micro/pkg/passwordhashing"
)

type UserService struct {
	userRepository    repository.UserRepository
	addressRepository repository.AddressRepository
	cnf               *configs.Config
}

func NewUserService(
	userRepository repository.UserRepository,
	addressRepository repository.AddressRepository,
	conf *configs.Config,
) *UserService {
	return &UserService{
		userRepository:    userRepository,
		addressRepository: addressRepository,
		cnf:               conf,
	}
}

func (service UserService) SignUp(signUpRequest dto.SignUpRequest) (dto.SignUpResponse, error) {
	// Hashing password from plain text
	userPassword := signUpRequest.Password
	hashedPassword, err := passwordhashing.HashPassword(userPassword, service.cnf.PasswordHash.Salt)
	if err != nil {
		return dto.SignUpResponse{}, err
	}
	signUpRequest.Password = hashedPassword

	user, err := service.userRepository.CreateUser(signUpRequest.ToModel())
	if err != nil {
		return dto.SignUpResponse{}, err
	}

	return dto.NewSignUpResponseFromEntity(user), nil
}

func (service UserService) SignIn(request dto.SignInRequest) (dto.SignInResponse, error) {
	user, err := service.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		return dto.SignInResponse{}, err
	}

	isPasswordMatched := passwordhashing.CheckPassword(request.Password,
		service.cnf.PasswordHash.Salt,
		user.Password)

	if !isPasswordMatched {
		return dto.SignInResponse{}, ce.ErrUnAuthorization
	}

	return dto.NewSignInResponse("", user), nil
}
