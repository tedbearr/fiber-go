package controller

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
	"github.com/tedbearr/go-learn/dto"
	"github.com/tedbearr/go-learn/helper"
	"github.com/tedbearr/go-learn/service"
)

type AuthController interface {
	Login(context *fiber.Ctx) error
	Register(context *fiber.Ctx) error
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (service *authController) Login(context *fiber.Ctx) error {
	var wg sync.WaitGroup
	var User dto.Login
	uniqueCode := helper.UniqueCode()

	if err := context.BodyParser(&User); err != nil {
		res := helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
		return context.Status(400).JSON(res)
	}

	slog.Info(uniqueCode+" Login request ", User)

	validate := helper.Validate(User)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return context.Status(200).JSON(res)
	}
	wg.Add(1)
	result, err := service.authService.Login(User, uniqueCode, &wg)
	wg.Wait()
	if err != nil {
		res := helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return context.Status(200).JSON(res)
	}

	res := helper.BuildResponse("00", "success", result)
	slog.Info(uniqueCode+" Login response ", res)
	return context.Status(200).JSON(res)
}

func (service *authController) Register(context *fiber.Ctx) error {
	var User dto.Register
	uniqueCode := helper.UniqueCode()

	if err := context.BodyParser(&User); err != nil {
		res := helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
		return context.Status(400).JSON(res)
	}

	slog.Info(uniqueCode+" Register request ", User)

	slog.Info(uniqueCode + " Register validate... ")
	validate := helper.Validate(User)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return context.Status(200).JSON(res)
	}

	register := service.authService.Register(User, uniqueCode)
	if register != nil {
		res := helper.BuildResponse("500", register.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return context.Status(200).JSON(res)
	}

	res := helper.BuildResponse("00", "success", helper.EmptyObj{})
	slog.Info(uniqueCode+" Register response ", res)
	return context.JSON(res)
}
