package route

import (
	"github.com/tedbearr/go-learn/controller"
	"github.com/tedbearr/go-learn/database"
	"github.com/tedbearr/go-learn/repository"
	"github.com/tedbearr/go-learn/service"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.DatabaseInit()

	authRepository repository.AuthRepository = repository.NewAuthRepository(db)
	authService    service.AuthService       = service.NewAuthService(authRepository)
	authController controller.AuthController = controller.NewAuthController(authService)

	globalParameterRepository repository.GlobalParameterRepository = repository.NewGlobalParameterRepository(db)
	globalParameterService    service.GlobalParameterService       = service.NewGlobalParameterService(globalParameterRepository)
	globalParameterController controller.GlobalParameterController = controller.NewGlobalParameterController(globalParameterService)
)
