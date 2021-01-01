package grpc

import (
	"context"
	cpro "github.com/Juno-chat-app/chatty-proto"
)

func (s *Server) CreateConnection(context.Context, *cpro.RequestMessage) (*cpro.ResponseMessage, error) {
	panic("not implemented")
}

func (s *Server) SendMessage(context.Context, *cpro.RequestMessage) (*cpro.ResponseMessage, error) {
	panic("not implemented")
}

func (s *Server) ReportChannelsMessages(context.Context, *cpro.RequestMessage) (*cpro.ResponseMessage, error) {
	panic("not implemented")
}

func (s *Server) FetchMessages(context.Context, *cpro.RequestMessage) (*cpro.ResponseMessage, error) {
	panic("not implemented")
}

func (s *Server) JoinConversation(context.Context, *cpro.RequestMessage) (*cpro.ResponseMessage, error) {
	panic("not implemented")
}

func (s *Server) AddMemberToConversation(context.Context, *cpro.RequestMessage) (*cpro.ResponseMessage, error) {
	panic("not implemented")
}
