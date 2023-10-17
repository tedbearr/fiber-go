package route

import (
	"github.com/gofiber/fiber/v2"
)

// var (
// 	db             *gorm.DB                  = database.DatabaseInit()
// 	authService    service.AuthService       = service.NewAuthService(db)
// 	authController controller.AuthController = controller.NewAuthController(authService)
// )

func AuthRoute(route fiber.Router) {
	group := route.Group("/auth")
	group.Post("/login", authController.Login)
	group.Post("/register", authController.Register)
}
