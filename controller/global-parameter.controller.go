package controller

import (
	"errors"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/tedbearr/go-learn/dto"
	"github.com/tedbearr/go-learn/helper"
	"github.com/tedbearr/go-learn/service"
	"gorm.io/gorm"
)

type GlobalParameterController interface {
	All(context *fiber.Ctx) error
	Insert(context *fiber.Ctx) error
	Find(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
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
	var globalParameter []dto.GlobalParameterAll = service.globalParameterService.All()
	result := helper.BuildResponse("00", "success get data", globalParameter)
	return context.Status(200).JSON(result)
}

func (service *globalParameterController) Insert(context *fiber.Ctx) error {
	var globalParameterCreate dto.GlobalParameterCreate
	var wg sync.WaitGroup

	if err := context.BodyParser(&globalParameterCreate); err != nil {
		res := helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
		return context.Status(400).JSON(res)
	}

	validate := helper.Validate(globalParameterCreate)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return context.Status(200).JSON(res)
	}
	wg.Add(1)
	insert := service.globalParameterService.Insert(globalParameterCreate, &wg)
	wg.Wait()
	if insert != nil {
		res := helper.BuildResponse("400", insert.Error(), helper.EmptyObj{})
		return context.Status(200).JSON(res)
	} else {
		res := helper.BuildResponse("00", "success", helper.EmptyObj{})
		return context.Status(200).JSON(res)
	}

}

func (service *globalParameterController) Find(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		return context.JSON(err.Error())
	}

	find, errFind := service.globalParameterService.Find(id)

	errors.Is(errFind, gorm.ErrRecordNotFound)

	if errFind != nil {
		res := helper.BuildResponse("400", errFind.Error(), helper.EmptyObj{})
		return context.JSON(res)
	}

	res := helper.BuildResponse("00", "success", find)
	return context.JSON(res)
}

func (service *globalParameterController) Update(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		return context.JSON(err.Error())
	}

	// _, errFind := service.globalParameterService.Find(id)

	// errors.Is(errFind, gorm.ErrRecordNotFound)

	// if errFind != nil {
	// 	res := helper.BuildResponse("400", errFind.Error(), helper.EmptyObj{})
	// 	return context.JSON(res)
	// }

	var globalParameterUpdate dto.GlobalParameterUpdate

	if err := context.BodyParser(&globalParameterUpdate); err != nil {
		res := helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
		return context.Status(400).JSON(res)
	}

	validate := helper.Validate(globalParameterUpdate)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return context.Status(200).JSON(res)
	}

	update := service.globalParameterService.Update(globalParameterUpdate, id)

	errors.Is(update, gorm.ErrRecordNotFound)

	if update != nil {
		res := helper.BuildResponse("400", update.Error(), helper.EmptyObj{})
		return context.JSON(res)
	} else {
		res := helper.BuildResponse("00", "success", helper.EmptyObj{})
		return context.JSON(res)
	}
}

func (service *globalParameterController) Delete(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		return context.JSON(err.Error())
	}

	update := service.globalParameterService.Update(dto.GlobalParameterUpdate{}, id)

	errors.Is(update, gorm.ErrRecordNotFound)

	if update != nil {
		res := helper.BuildResponse("400", update.Error(), helper.EmptyObj{})
		return context.JSON(res)
	} else {
		res := helper.BuildResponse("00", "success", helper.EmptyObj{})
		return context.JSON(res)
	}
}
