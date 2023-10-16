package service

import (
	"github.com/tedbearr/go-learn/entity"
	"gorm.io/gorm"
)

type GlobalParameterService interface {
	All() []entity.GlobalParameter
}

type globalParameterConnection struct {
	connection *gorm.DB
}

func NewGlobalParameterService(dbConnection *gorm.DB) GlobalParameterService {
	return &globalParameterConnection{
		connection: dbConnection,
	}
}

func (db *globalParameterConnection) All() []entity.GlobalParameter {
	var globalParameter []entity.GlobalParameter
	db.connection.Table("global_parameter").Find(&globalParameter)
	return globalParameter
}
