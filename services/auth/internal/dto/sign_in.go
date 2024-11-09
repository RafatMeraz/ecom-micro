package dto

import "github.com/RafatMeraz/ecom-micro/auth/internal/model"

type SignInRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignInResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func NewSignInResponse(token string, user model.User) SignInResponse {
	return SignInResponse{
		Token: token,
		User:  NewUserResponseFromEntity(user),
	}
}
