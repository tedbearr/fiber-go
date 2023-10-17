package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedbearr/go-learn/middleware"
)

// var (
// 	// dbConnection              *gorm.DB                             = database.DatabaseInit()
// 	globalParameterService    service.GlobalParameterService       = service.NewGlobalParameterService(db)
// 	globalParameterController controller.GlobalParameterController = controller.NewGlobalParameterController(globalParameterService)
// )

func GlobalParameterRoute(route fiber.Router) {
	group := route.Group("/global-parameter", middleware.Jwt)
	group.Get("/", globalParameterController.All)
	group.Get("/:id", globalParameterController.Find)
	group.Post("/insert", globalParameterController.Insert)
	group.Post("update/:id", globalParameterController.Update)
	group.Post("/delete/:id", globalParameterController.Delete)
}
