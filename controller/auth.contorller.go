package controller

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
	"github.com/tedbearr/go-learn/dto"
	"github.com/tedbearr/go-learn/helper"
	"github.com/tedbearr/go-learn/service"
	"gorm.io/gorm"
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

	slog.Info(uniqueCode + " Login check auth... ")
	user, errCheck := service.authService.CheckUser(User.Username)
	errors.Is(errCheck, gorm.ErrRecordNotFound)
	if errCheck != nil {
		res := helper.BuildResponse("400", errCheck.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return context.Status(400).JSON(res)
	}

	slog.Info(uniqueCode + " Login compare password... ")
	comparePassword := service.authService.ComparePassword(user.Password, []byte(User.Password))
	if comparePassword != nil {
		res := helper.BuildResponse("400", comparePassword.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return context.Status(400).JSON(res)
	}

	slog.Info(uniqueCode + " Login generating access token... ")
	accessToken, errAccessToken := service.authService.GenerateAccessToken()
	if errAccessToken != nil {
		res := helper.BuildResponse("400", errAccessToken.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return context.Status(400).JSON(res)
	}

	slog.Info(uniqueCode + " Login generating refresh token... ")
	refreshToken, errRefreshToken := service.authService.GenerateRefreshToken()
	if errRefreshToken != nil {
		res := helper.BuildResponse("400", errRefreshToken.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return context.Status(400).JSON(res)
	}

	updateTokenData := dto.Auth{
		RefreshToken: refreshToken,
	}

	slog.Info(uniqueCode + " Login updating refresh token... ")
	updateToken := service.authService.UpdateRefreshToken(updateTokenData, User.Username)
	if updateToken != nil {
		res := helper.BuildResponse("400", updateToken.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return context.Status(400).JSON(res)
	}

	data := dto.ResponseToken{AccessToken: accessToken, RefreshToken: refreshToken}

	res := helper.BuildResponse("00", "success", data)
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

	slog.Info(uniqueCode + " Register check user... ")
	_, errCheck := service.authService.CheckUser(User.Username)
	errors.Is(errCheck, gorm.ErrDuplicatedKey)

	slog.Info(uniqueCode + " Register hashing password... ")
	hashedPassword, errHash := service.authService.HashPassword(User.Password)
	if errHash != nil {
		res := helper.BuildResponse("400", errHash.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return context.JSON(res)
	}

	dataInsert := dto.Auth{
		Username:  User.Username,
		Email:     User.Email,
		Password:  hashedPassword,
		Name:      User.Name,
		StatusID:  1,
		CreatedAt: time.Now(),
	}

	slog.Info(uniqueCode + " Register insert data... ")
	insert := service.authService.Insert(dataInsert)
	if insert != nil {
		res := helper.BuildResponse("400", insert.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return context.JSON(res)
	}

	res := helper.BuildResponse("00", "success", helper.EmptyObj{})
	slog.Info(uniqueCode+" Register response ", res)
	return context.JSON(res)
}
