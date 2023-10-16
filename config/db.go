package config

import "github.com/tedbearr/go-learn/database"

func DatabaseConfig() {
	// database.DatabaseInit()
	gorm := database.DatabaseInit()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}
	dbGorm.Ping()
}
