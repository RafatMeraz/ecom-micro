package handler

import (
	"github.com/RafatMeraz/ecom-micro/auth/internal/dto"
	"github.com/RafatMeraz/ecom-micro/auth/internal/service"
	ce "github.com/RafatMeraz/ecom-micro/pkg/errors"
	"github.com/RafatMeraz/ecom-micro/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct {
	userService *service.UserService
	validator   *validator.Validate
}

func NewAuthHandler(
	userService *service.UserService,
	validator *validator.Validate) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		validator:   validator,
	}
}

func (a AuthHandler) SignUp(c *fiber.Ctx) error {
	var signUpRequest dto.SignUpRequest
	if err := c.BodyParser(&signUpRequest); err != nil {
		statusCode, res := ce.GetErrorResponse(ce.ErrInvalidData)
		return c.Status(statusCode).JSON(res)
	}

	if err := a.validator.Struct(signUpRequest); err != nil {
		statusCode, res := ce.GetErrorResponse(ce.ErrRequiredFieldsMissing)
		return c.Status(statusCode).JSON(res)
	}

	signUpRes, err := a.userService.SignUp(signUpRequest)
	if err != nil {
		statusCode, res := ce.GetErrorResponse(err)
		return c.Status(statusCode).JSON(res)
	}

	_, res := response.GenerateResponse(signUpRes, nil, false)

	return c.Status(http.StatusCreated).JSON(res)
}

func (a AuthHandler) SignIn(c *fiber.Ctx) error {
	var signInRequest dto.SignInRequest
	if err := c.BodyParser(&signInRequest); err != nil {
		statusCode, res := ce.GetErrorResponse(ce.ErrInvalidData)
		return c.Status(statusCode).JSON(res)
	}

	if err := a.validator.Struct(signInRequest); err != nil {
		statusCode, res := ce.GetErrorResponse(ce.ErrRequiredFieldsMissing)
		return c.Status(statusCode).JSON(res)
	}
	signInRes, err := a.userService.SignIn(signInRequest)
	if err != nil {
		statusCode, res := ce.GetErrorResponse(err)
		return c.Status(statusCode).JSON(res)
	}
	statusCode, res := response.GenerateResponse(signInRes, nil, false)
	return c.Status(statusCode).JSON(res)
}
