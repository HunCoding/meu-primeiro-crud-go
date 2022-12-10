package view

import (
	"github.com/HunCoding/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/HunCoding/meu-primeiro-crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
