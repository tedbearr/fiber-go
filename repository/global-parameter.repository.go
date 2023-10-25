package repository

import (
	"github.com/tedbearr/go-learn/dto"
	"gorm.io/gorm"
)

type GlobalParameterRepository interface {
	All() []dto.GlobalParameterAll
	Insert(dto.GlobalParameter) error
	Find(id int) (dto.GlobalParameter, error)
	Update(globalParameter dto.GlobalParameter, id int) error
	Delete(globalParameter dto.GlobalParameter, id int) error
	Count() (int, error)
}

type globalParameterRepository struct {
	connection *gorm.DB
}

func NewGlobalParameterRepository(dbConnection *gorm.DB) GlobalParameterRepository {
	return &globalParameterRepository{
		connection: dbConnection,
	}
}

func (db *globalParameterRepository) All() []dto.GlobalParameterAll {
	var globalParameter []dto.GlobalParameterAll
	db.connection.Table("global_parameter").
		Select("global_parameter.name, global_parameter.code, global_parameter.value, status.name as status_id").
		Joins("left join status on status.id = global_parameter.status_id").
		Where("global_parameter.status_id = ?", 1).
		Find(&globalParameter)
	return globalParameter
}

func (db *globalParameterRepository) Insert(globalParameter dto.GlobalParameter) error {
	err := db.connection.Table("global_parameter").
		Create(&globalParameter).Error
	return err
}

func (db *globalParameterRepository) Find(id int) (dto.GlobalParameter, error) {
	var globalParameter dto.GlobalParameter
	check := db.connection.Table("global_parameter").
		First(&globalParameter, id).Error
	return globalParameter, check
}

func (db *globalParameterRepository) Update(globalParameter dto.GlobalParameter, id int) error {
	err := db.connection.Table("global_parameter").Where("id = ?", id).
		Updates(globalParameter).Error
	return err
}

func (db *globalParameterRepository) Delete(globalParameter dto.GlobalParameter, id int) error {
	err := db.connection.Table("global_parameter").Where("id = ?", id).
		Updates(globalParameter).Error
	return err
}

func (db *globalParameterRepository) Count() (int, error) {
	// defer wg.Done()
	var globalParameter []dto.GlobalParameterAll
	// wg.Add(1)
	res := db.connection.Table("global_parameter").Find(&globalParameter)
	// wg.Done()
	// wg.Wait()
	// fmt.Println(len(globalParameter))
	return len(globalParameter), res.Error
}
