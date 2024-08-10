package usecase

import (
	"context"
	"fmt"
	"user-svc/internal/dto"
	"user-svc/pkg/http_client"
	"user-svc/pkg/producer"
)

type Usecase struct {
	userClient        *http_client.UserClient
	orchestraProducer *producer.KafkaProducer
}

func NewUsecase(userClient *http_client.UserClient, orchestraProducer *producer.KafkaProducer) *Usecase {
	return &Usecase{
		userClient:        userClient,
		orchestraProducer: orchestraProducer,
	}
}

func (u *Usecase) ValidateUser(ctx context.Context, request *dto.UserValidateRequest) (*dto.UserResponse, error) {

	var response dto.BaseResponse[dto.UserResponse]

	err := u.userClient.GET(
		fmt.Sprintf("/users/%s", request.UserID),
		nil,
		&response,
	)

	if err != nil {
		return nil, err
	}

	// u.orchestraProducer.SendMessage()
	return &response.Data, nil
}
