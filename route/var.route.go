package route

import (
	"github.com/tedbearr/go-learn/controller"
	"github.com/tedbearr/go-learn/database"
	"github.com/tedbearr/go-learn/service"
	"gorm.io/gorm"
)

var (
	db                        *gorm.DB                             = database.DatabaseInit()
	authService               service.AuthService                  = service.NewAuthService(db)
	authController            controller.AuthController            = controller.NewAuthController(authService)
	globalParameterService    service.GlobalParameterService       = service.NewGlobalParameterService(db)
	globalParameterController controller.GlobalParameterController = controller.NewGlobalParameterController(globalParameterService)
)
