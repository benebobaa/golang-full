package usecase

import (
	"user-svc/internal/dto"
	"user-svc/pkg/http_client"
)

type Usecase struct {
	userClient *http_client.UserClient
}

func NewUsecase(userClient *http_client.UserClient) *Usecase {
	return &Usecase{userClient: userClient}
}

func (u *Usecase) ValidateUser(request *dto.UserValidateRequest) (*dto.UserValidateResponse, error) {

	var response dto.UserValidateResponse

	err := u.userClient.Call("POST", request, &response)

	if err != nil {
		return nil, err
	}

	response.Status = "success"
	response.Message = "User validated successfully"

	return &response, nil
}
