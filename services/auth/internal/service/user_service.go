package service

import (
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/RafatMeraz/ecom-micro/auth/internal/dto"
	"github.com/RafatMeraz/ecom-micro/auth/internal/repository"
	ce "github.com/RafatMeraz/ecom-micro/pkg/errors"
	"github.com/RafatMeraz/ecom-micro/pkg/passwordhashing"
	"github.com/RafatMeraz/ecom-micro/pkg/service"
)

type UserService struct {
	userRepository    repository.UserRepository
	addressRepository repository.AddressRepository
	cnf               *configs.Config
	jwtService        *service.JwtService
}

func NewUserService(
	userRepository repository.UserRepository,
	addressRepository repository.AddressRepository,
	conf *configs.Config,
	jwtService *service.JwtService,
) *UserService {
	return &UserService{
		userRepository:    userRepository,
		addressRepository: addressRepository,
		cnf:               conf,
		jwtService:        jwtService,
	}
}

func (s UserService) SignUp(signUpRequest dto.SignUpRequest) (dto.SignUpResponse, error) {
	// Hashing password from plain text
	userPassword := signUpRequest.Password
	hashedPassword, err := passwordhashing.HashPassword(userPassword, s.cnf.PasswordHash.Salt)
	if err != nil {
		return dto.SignUpResponse{}, err
	}
	signUpRequest.Password = hashedPassword

	user, err := s.userRepository.CreateUser(signUpRequest.ToModel())
	if err != nil {
		return dto.SignUpResponse{}, err
	}

	tokens, err := s.jwtService.PrepareJwtTokens(user.ID)
	if err != nil {
		return dto.SignUpResponse{}, err
	}

	return dto.NewSignUpResponseFromEntity(tokens, user), nil
}

func (s UserService) SignIn(request dto.SignInRequest) (dto.SignInResponse, error) {
	user, err := s.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		return dto.SignInResponse{}, err
	}

	isPasswordMatched := passwordhashing.CheckPassword(request.Password,
		s.cnf.PasswordHash.Salt,
		user.Password)

	if !isPasswordMatched {
		return dto.SignInResponse{}, ce.ErrUnAuthorization
	}

	tokens, err := s.jwtService.PrepareJwtTokens(user.ID)
	if err != nil {
		return dto.SignInResponse{}, err
	}

	return dto.NewSignInResponse(tokens, user), nil
}
