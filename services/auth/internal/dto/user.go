package dto

import "github.com/RafatMeraz/ecom-micro/auth/internal/model"

type UserResponse struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewUserResponseFromEntity(u model.User) UserResponse {
	return UserResponse{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
