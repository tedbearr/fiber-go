package service

import (
	"errors"
	"time"

	"github.com/tedbearr/go-learn/dto"
	"github.com/tedbearr/go-learn/repository"
	"gorm.io/gorm"
)

type GlobalParameterService interface {
	All() []dto.GlobalParameterAll
	Insert(data dto.GlobalParameterCreate) error
	Find(id int) (dto.GlobalParameter, error)
	Update(globalParameter dto.GlobalParameterUpdate, id int) error
	Delete(globalParameter dto.GlobalParameterUpdate, id int) error
}

type globalParameterService struct {
	connection repository.GlobalParameterRepository
}

func NewGlobalParameterService(repository repository.GlobalParameterRepository) GlobalParameterService {
	return &globalParameterService{
		connection: repository,
	}
}

func (repository *globalParameterService) All() []dto.GlobalParameterAll {
	return repository.connection.All()
}

func (repository *globalParameterService) Insert(globalParameterCreate dto.GlobalParameterCreate) error {
	dataInsert := dto.GlobalParameter{
		Code:      globalParameterCreate.Code,
		Value:     globalParameterCreate.Value,
		Name:      globalParameterCreate.Name,
		StatusID:  1,
		CreatedAt: time.Now(),
	}
	return repository.connection.Insert(dataInsert)
}

func (repository *globalParameterService) Find(id int) (dto.GlobalParameter, error) {
	return repository.connection.Find(id)
}

func (repository *globalParameterService) Update(globalParameterUpdate dto.GlobalParameterUpdate, id int) error {
	_, errFind := repository.connection.Find(id)

	errors.Is(errFind, gorm.ErrRecordNotFound)

	if errFind != nil {
		return errFind
	}

	dataInsert := dto.GlobalParameter{
		Name:      globalParameterUpdate.Name,
		Code:      globalParameterUpdate.Code,
		Value:     globalParameterUpdate.Value,
		UpdatedAt: time.Now(),
	}
	return repository.connection.Update(dataInsert, id)
}

func (repository *globalParameterService) Delete(globalParameter dto.GlobalParameterUpdate, id int) error {
	_, errFind := repository.connection.Find(id)

	errors.Is(errFind, gorm.ErrRecordNotFound)

	if errFind != nil {
		return errFind
	}

	dataInsert := dto.GlobalParameter{
		StatusID:  2,
		UpdatedAt: time.Now(),
	}
	return repository.connection.Delete(dataInsert, id)
}
