package handler

import (
	"context"
	"google.golang.org/grpc"
	"userservice/pkg/v1/user/service"
	"userservice/proto/user"
)

type UserHandler struct {
	UserService service.UserServiceImpl
	user.UnimplementedUserServiceServer
}
type UserHandlerImpl interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error)
	DeactivateUser(ctx context.Context, req *user.DeactivateUserRequest) (*user.DeactivateUserResponse, error)
}

func NewUserHandler(grpcServer *grpc.Server, service service.UserServiceImpl) UserHandlerImpl {
	userGrpc := &UserHandler{UserService: service}
	user.RegisterUserServiceServer(grpcServer, userGrpc)
	return userGrpc
}

func (srv *UserHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return &user.CreateUserResponse{}, nil
}

func (srv *UserHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	return &user.UpdateUserResponse{}, nil
}

func (srv *UserHandler) DeactivateUser(ctx context.Context, req *user.DeactivateUserRequest) (*user.DeactivateUserResponse, error) {
	return &user.DeactivateUserResponse{}, nil
}
