package service_test

import (
	"testing"

	"github.com/tedbearr/go-learn/repository"
	"github.com/tedbearr/go-learn/service"
	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	globalParameterRepo repository.GlobalParameterRepository = repository.NewGlobalParameterRepository(DB)
	globalParameter     service.GlobalParameterService       = service.NewGlobalParameterService(globalParameterRepo)
	correctResponse     error                                = nil
)

func AllTest(t *testing.T) {
	t.Logf("testing")

	_, err := globalParameter.Find(1)

	if err != correctResponse {
		t.Errorf("SALAH! harusnya %.2f", correctResponse)
	}

	// if globalParameter.All() != correctResponse {

	// }
}
