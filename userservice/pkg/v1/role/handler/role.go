package handler

import (
	"context"
	"google.golang.org/grpc"
	"userservice/pkg/v1/role/service"
	"userservice/proto/role"
)

type RoleHandler struct {
	RoleService service.RoleServiceImpl
	role.UnimplementedRoleServiceServer
}

type RoleHandlerImpl interface {
	UpdateRole(ctx context.Context, req *role.UpdateRoleRequest) (*role.UpdateRoleResponse, error)
	GetAllRoles(ctx context.Context, req *role.GetAllRolesRequest) (*role.GetAllRolesResponse, error)
	GetRole(ctx context.Context, req *role.GetRoleRequest) (*role.GetRoleResponse, error)
	CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.CreateRoleResponse, error)
	DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.DeleteRoleResponse, error)
}

func NewRoleHandler(grpcServer *grpc.Server, service service.RoleServiceImpl) RoleHandlerImpl {
	permissionGrpc := &RoleHandler{RoleService: service}
	role.RegisterRoleServiceServer(grpcServer, permissionGrpc)
	return permissionGrpc
}

func (srv *RoleHandler) UpdateRole(ctx context.Context, req *role.UpdateRoleRequest) (*role.UpdateRoleResponse, error) {
	return &role.UpdateRoleResponse{}, nil
}

func (srv *RoleHandler) GetAllRoles(ctx context.Context, req *role.GetAllRolesRequest) (*role.GetAllRolesResponse, error) {
	return &role.GetAllRolesResponse{}, nil
}

func (srv *RoleHandler) GetRole(ctx context.Context, req *role.GetRoleRequest) (*role.GetRoleResponse, error) {
	return &role.GetRoleResponse{}, nil
}

func (srv *RoleHandler) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.CreateRoleResponse, error) {
	return &role.CreateRoleResponse{}, nil
}

func (srv *RoleHandler) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.DeleteRoleResponse, error) {
	return &role.DeleteRoleResponse{}, nil
}
