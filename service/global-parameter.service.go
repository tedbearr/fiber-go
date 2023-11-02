package service

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/tedbearr/go-learn/dto"
	"github.com/tedbearr/go-learn/helper"
	"github.com/tedbearr/go-learn/repository"
	"gorm.io/gorm"
)

type GlobalParameterService interface {
	All() []dto.GlobalParameterAll
	Insert(data dto.GlobalParameterCreate, wg *sync.WaitGroup) error
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
	// var res []dto.GlobalParameterAll
	var messages = make(chan []dto.GlobalParameterAll)
	go func() {
		messages <- repository.connection.All()
	}()

	return <-messages
}

func (repository *globalParameterService) Insert(globalParameterCreate dto.GlobalParameterCreate, wg *sync.WaitGroup) error {
	defer wg.Done()
	var m sync.Mutex
	m.Lock()
	wg.Add(1)
	a, _ := repository.connection.Count(wg)
	num := a + 1
	count := strconv.Itoa(num)
	lpad := helper.Lpad(count, "0", 3)
	code := "GP" + lpad

	// fmt.Println(code)
	// var code = make(chan string)
	// wg.Add(1)
	// go func() {
	// defer wg.Done()
	// code <- repository.sequenceCodeGlobalParameter(wg)
	// }()
	// wg.Wait()
	dataInsert := dto.GlobalParameter{
		Code:      code,
		Value:     code,
		Name:      code,
		StatusID:  1,
		CreatedAt: time.Now(),
	}

	insert := repository.connection.Insert(dataInsert)
	defer m.Unlock()
	return insert
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

// func (repository *globalParameterService) sequenceCodeGlobalParameter() string {
// 	// wg.Add(1)
// 	var wg sync.WaitGroup
// 	var m sync.Mutex
// 	wg.Add(1)
// 	m.Lock()
// 	count, err := repository.connection.Count(&wg)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	m.Unlock()
// 	wg.Wait()
// 	num := count + 1
// 	res := strconv.Itoa(num)
// 	// fmt.Println(count)
// 	return res
// }
