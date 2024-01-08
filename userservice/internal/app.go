package internal

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"userservice/config"
	"userservice/internal/database"
	"userservice/internal/models"
	permissionHandler "userservice/pkg/v1/permission/handler"
	permissionRepository "userservice/pkg/v1/permission/repository"
	permissionService "userservice/pkg/v1/permission/service"
	roleHandler "userservice/pkg/v1/role/handler"
	roleRepository "userservice/pkg/v1/role/repository"
	roleService "userservice/pkg/v1/role/service"
	userHandler "userservice/pkg/v1/user/handler"
	userRepository "userservice/pkg/v1/user/repository"
	userService "userservice/pkg/v1/user/service"
)

type Application struct {
	DB                *gorm.DB
	Logger            *logrus.Logger
	UserHandler       userHandler.UserHandlerImpl
	PermissionHandler permissionHandler.PermissionHandlerImpl
	RoleHandler       roleHandler.RoleHandlerImpl
	grpcServer        *grpc.Server
}

// Initialize initializes the application and its components.
func Initialize(server *grpc.Server) *Application {
	//Load environment variables
	variables, err := config.Load()
	if err != nil {
		log.Printf("Unable to read enviroment variables %s", err)
	}
	// Initialize the database
	db, err := database.InitDB(variables)
	if err != nil {
		log.Printf("Unable to connect to the database %s", err)
	}
	//Initialize logrus logger
	logger := logrus.New()
	//Initialize Services
	userSvc := userService.NewUserService(userRepository.NewUserRepository(db))
	roleSvc := roleService.NewRoleService(roleRepository.NewUserRepository(db))
	permissionSvc := permissionService.NewPermissionService(permissionRepository.NewPermissionRepository(db))
	//Initialize roles
	InitializeRoles(db)

	// Return the initialized application
	return &Application{
		DB:                db,
		Logger:            logger,
		UserHandler:       userHandler.NewUserHandler(server, userSvc),
		PermissionHandler: permissionHandler.NewPermissionHandler(server, roleSvc),
		RoleHandler:       roleHandler.NewRoleHandler(server, permissionSvc),
		grpcServer:        server,
	}
}

func InitializeRoles(db *gorm.DB) {
	defaultRoles := []models.Role{
		{
			Name: models.RoleAdmin,
			Permissions: []models.Permission{
				{Name: "manage_users", Description: "Manage users"},
				{Name: "manage_roles", Description: "Manage roles"},
				// Add more permissions as needed
			},
		},
		{
			Name: models.RoleSuperAdmin,
			Permissions: []models.Permission{
				{Name: "manage_everything", Description: "Manage everything"},
				// Add more permissions as needed
			},
		},
		{
			Name: models.RoleCustomer,
			Permissions: []models.Permission{
				{Name: "view_profile", Description: "View user profile"},
				// Add more permissions as needed
			},
		},
	}

	for _, role := range defaultRoles {
		if err := db.Create(&role).Error; err != nil {
			log.Fatalf("Error creating role %s: %v", role.Name, err)
		}
	}
}
