package user

import (
	"context"
	"github.com/Juno-chat-app/chatty-service/infra/logger"
)

type GetUserRequest struct {
	Email       string
	PhoneNumber string
}

type User struct {
	Email       string
	PhoneNumber string
	UserId      string
	UserName    string
}

type IUserService interface {
	GetUser(ctx context.Context, request GetUserRequest) (*User, error)
}

func NewUserServiceClient(logger logger.ILogger, host string, port int32) IUserService {
	service := iUserService{
		host:   host,
		port:   port,
		logger: logger,
		client: nil,
	}

	return &service
}
