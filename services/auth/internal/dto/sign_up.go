package dto

import "github.com/RafatMeraz/ecom-micro/auth/internal/model"

type SignUpRequest struct {
	FirstName string `json:"first_name" form:"first_name" validate:"required"`
	LastName  string `json:"last_name" form:"last_name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required"`
}

func (r SignUpRequest) ToModel() model.User {
	return model.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
	}
}

type SignUpResponse struct {
	Token        map[string]any `json:"token"`
	UserResponse UserResponse   `json:"user"`
}

func NewSignUpResponseFromEntity(token map[string]any, user model.User) SignUpResponse {
	return SignUpResponse{
		Token:        token,
		UserResponse: NewUserResponseFromEntity(user),
	}
}
