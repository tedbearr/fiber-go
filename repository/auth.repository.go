package repository

import (
	"github.com/tedbearr/go-learn/dto"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Insert(authData *dto.Auth) error
	CheckUsername(username string) (dto.Auth, error)
	UpdateRefreshToken(authData *dto.Auth, username string) error
	CheckEmail(email string) (dto.Auth, error)
}

type authRepository struct {
	connection *gorm.DB
}

func NewAuthRepository(dbConnection *gorm.DB) AuthRepository {
	return &authRepository{
		connection: dbConnection,
	}
}

func (db *authRepository) Insert(authData *dto.Auth) error {
	err := db.connection.Table("users").
		Create(&authData).Error
	return err
}

func (db *authRepository) CheckUsername(username string) (dto.Auth, error) {
	var auth dto.Auth
	check := db.connection.Table("users").
		Where("username = ?", username).
		First(&auth).Error
	return auth, check
}

func (db *authRepository) CheckEmail(email string) (dto.Auth, error) {
	var auth dto.Auth
	check := db.connection.Table("users").
		Where("email = ?", email).
		First(&auth).Error
	return auth, check
}

func (db *authRepository) UpdateRefreshToken(authData *dto.Auth, username string) error {
	update := db.connection.Table("users").
		Where("username = ?", username).
		Updates(&authData).Error
	return update
}
