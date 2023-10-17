package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tedbearr/go-learn/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	CheckUser(username string) (dto.Auth, error)
	GenerateAccessToken() (string, error)
	GenerateRefreshToken() (string, error)
	Insert(dto.Auth) error
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, plainPassword []byte) error
	UpdateRefreshToken(authData dto.Auth, username string) error
}

type authConnection struct {
	connection *gorm.DB
}

func NewAuthService(dbConnection *gorm.DB) AuthService {
	return &authConnection{
		connection: dbConnection,
	}
}

func (db *authConnection) CheckUser(username string) (dto.Auth, error) {
	var auth dto.Auth
	check := db.connection.Table("users").
		Where("username = ?", username).
		First(&auth).Error
	return auth, check
}

func (db *authConnection) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

func (db *authConnection) Insert(authData dto.Auth) error {
	err := db.connection.Table("users").
		Create(&authData).Error
	return err
}

func (db *authConnection) GenerateAccessToken() (string, error) {
	tokenSecret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	claims := jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resToken, err := token.SignedString([]byte(tokenSecret))

	return resToken, err
}

func (db *authConnection) GenerateRefreshToken() (string, error) {
	tokenSecret := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	claims := jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 10000).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resToken, err := token.SignedString([]byte(tokenSecret))

	return resToken, err
}

func (db *authConnection) ComparePassword(hashedPassword string, plainPassword []byte) error {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err
}

func (db *authConnection) UpdateRefreshToken(authData dto.Auth, username string) error {
	update := db.connection.Table("users").
		Where("username = ?", username).
		Updates(&authData).Error
	return update
}
