package user

import (
	"context"
	"fmt"
	"github.com/Juno-chat-app/chatty-service/infra/logger"
	usrpro "github.com/Juno-chat-app/user-proto"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type iUserService struct {
	host   string
	port   int32
	logger logger.ILogger
	client usrpro.UserServiceClient
}

func (i *iUserService) GetUser(ctx context.Context, request GetUserRequest) (*User, error) {
	if i.client == nil {
		err := i.connect()
		if err != nil {
			return nil, err
		}
	}

	requestBody := usrpro.GetUserRequest{
		SearchInfo: &usrpro.GetUserRequest_ConnectionInfo{
			PhoneNumber: request.PhoneNumber,
			Email:       request.Email,
		},
	}
	requestBodyByte, err := proto.Marshal(&requestBody)
	if err != nil {
		return nil, err
	}

	outgoingRequest := usrpro.RequestMessage{
		Name: "GetUserRequest",
		Type: "data",
		Time: time.Now().UTC().String(),
		Body: &usrpro.Any{
			TypeUrl: "GetUserRequest",
			Value:   requestBodyByte,
		},
	}

	response, err := i.client.GetUser(ctx, &outgoingRequest)
	if err != nil {
		return nil, err
	}

	responseBody := usrpro.GetUserResponse{}
	err = proto.Unmarshal(response.Data.Value, &responseBody)
	if err != nil {
		i.logger.Error(err.Error(),
			"method", "iUserService.GetUser.Unmarshal")
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}

	user := User{
		Email:       responseBody.Info.Email,
		PhoneNumber: responseBody.Info.PhoneNumber,
		UserId:      responseBody.Info.UserId,
		UserName:    responseBody.Info.UserName,
	}

	return &user, nil
}

func (i *iUserService) connect() error {
	address := fmt.Sprintf("%s:%d", i.host, i.port)
	connection, err := grpc.Dial(address)
	if err != nil {
		i.logger.Error(err.Error(),
			"method", "iUserService.connect",
			"host", i.host,
			"port", i.port)
		return status.Error(http.StatusInternalServerError, err.Error())
	}

	i.client = usrpro.NewUserServiceClient(connection)
	return nil
}
