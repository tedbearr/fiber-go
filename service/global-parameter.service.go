package service

import (
	"github.com/tedbearr/go-learn/dto"
	"gorm.io/gorm"
)

type GlobalParameterService interface {
	All() []dto.GlobalParameterAll
	Insert(dto.GlobalParameter) error
	Find(id int) (dto.GlobalParameter, error)
	Update(globalParameter dto.GlobalParameter, id int) error
	Delete(globalParameter dto.GlobalParameter, id int) error
}

type globalParameterConnection struct {
	connection *gorm.DB
}

func NewGlobalParameterService(dbConnection *gorm.DB) GlobalParameterService {
	return &globalParameterConnection{
		connection: dbConnection,
	}
}

func (db *globalParameterConnection) All() []dto.GlobalParameterAll {
	var globalParameter []dto.GlobalParameterAll
	db.connection.Table("global_parameter").
		Select("global_parameter.name, global_parameter.code, global_parameter.value, status.name as status_id").
		Joins("left join status on status.id = global_parameter.status_id").
		Where("global_parameter.status_id = ?", 1).
		Find(&globalParameter)
	return globalParameter
}

func (db *globalParameterConnection) Insert(globalParameter dto.GlobalParameter) error {
	err := db.connection.Table("global_parameter").
		Create(&globalParameter).Error
	return err
}

func (db *globalParameterConnection) Find(id int) (dto.GlobalParameter, error) {
	var globalParameter dto.GlobalParameter
	check := db.connection.Table("global_parameter").
		First(&globalParameter, id).Error
	return globalParameter, check
}

func (db *globalParameterConnection) Update(globalParameter dto.GlobalParameter, id int) error {
	err := db.connection.Table("global_parameter").Where("id = ?", id).
		Updates(globalParameter).Error
	return err
}

func (db *globalParameterConnection) Delete(globalParameter dto.GlobalParameter, id int) error {
	err := db.connection.Table("global_parameter").Where("id = ?", id).
		Updates(globalParameter).Error
	return err
}
