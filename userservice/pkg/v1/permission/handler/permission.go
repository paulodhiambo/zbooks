package handler

import (
	"context"
	"google.golang.org/grpc"
	"userservice/pkg/v1/permission/service"
	"userservice/proto/permission"
)

type PermissionHandler struct {
	PermissionService service.PermissionServiceImpl
	permission.UnimplementedPermissionServiceServer
}

type PermissionHandlerImpl interface {
	CreatePermission(ctx context.Context, req *permission.CreatePermissionRequest) (*permission.CreatePermissionResponse, error)
	GetPermission(ctx context.Context, req *permission.GetPermissionRequest) (*permission.GetPermissionResponse, error)
	GetAllPermissions(ctx context.Context, req *permission.GetAllPermissionsRequest) (*permission.GetAllPermissionsResponse, error)
	UpdatePermission(ctx context.Context, req *permission.UpdatePermissionRequest) (*permission.UpdatePermissionResponse, error)
	DeletePermission(ctx context.Context, req *permission.DeletePermissionRequest) (*permission.DeletePermissionResponse, error)
}

func NewPermissionHandler(grpcServer *grpc.Server, service service.PermissionServiceImpl) PermissionHandlerImpl {
	permissionGrpc := &PermissionHandler{PermissionService: service}
	permission.RegisterPermissionServiceServer(grpcServer, permissionGrpc)
	return permissionGrpc
}

func (srv *PermissionHandler) CreatePermission(ctx context.Context, req *permission.CreatePermissionRequest) (*permission.CreatePermissionResponse, error) {
	return &permission.CreatePermissionResponse{}, nil
}

func (srv *PermissionHandler) GetPermission(ctx context.Context, req *permission.GetPermissionRequest) (*permission.GetPermissionResponse, error) {
	return &permission.GetPermissionResponse{}, nil
}

func (srv *PermissionHandler) GetAllPermissions(ctx context.Context, req *permission.GetAllPermissionsRequest) (*permission.GetAllPermissionsResponse, error) {
	return &permission.GetAllPermissionsResponse{}, nil
}

func (srv *PermissionHandler) UpdatePermission(ctx context.Context, req *permission.UpdatePermissionRequest) (*permission.UpdatePermissionResponse, error) {
	return &permission.UpdatePermissionResponse{}, nil
}

func (srv *PermissionHandler) DeletePermission(ctx context.Context, req *permission.DeletePermissionRequest) (*permission.DeletePermissionResponse, error) {
	return &permission.DeletePermissionResponse{}, nil
}
