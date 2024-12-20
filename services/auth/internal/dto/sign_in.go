package dto

import "github.com/RafatMeraz/ecom-micro/auth/internal/model"

type SignInRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignInResponse struct {
	Token map[string]any `json:"token"`
	User  UserResponse   `json:"user"`
}

func NewSignInResponse(token map[string]any, user model.User) SignInResponse {
	return SignInResponse{
		Token: token,
		User:  NewUserResponseFromEntity(user),
	}
}
