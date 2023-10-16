package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedbearr/go-learn/entity"
	"github.com/tedbearr/go-learn/service"
)

type GlobalParameterController interface {
	All(context *fiber.Ctx) error
}

type globalParameterController struct {
	globalParameterService service.GlobalParameterService
}

func NewGlobalParameterController(globalParameterService service.GlobalParameterService) GlobalParameterController {
	return &globalParameterController{
		globalParameterService: globalParameterService,
	}
}

func (service *globalParameterController) All(context *fiber.Ctx) error {
	var globalParameter []entity.GlobalParameter = service.globalParameterService.All()
	return context.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to Golang, Fiber, and GORM",
		"data":    globalParameter,
	})
}
