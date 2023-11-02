package service

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gookit/slog"
	"github.com/tedbearr/go-learn/dto"
	"github.com/tedbearr/go-learn/helper"
	"github.com/tedbearr/go-learn/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(data dto.Login, uniqueCode string, wg *sync.WaitGroup) (interface{}, error)
	Register(data dto.Register, uniqueCode string) error
}

type authService struct {
	connection repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		connection: authRepository,
	}
}

func (repository *authService) Login(authData dto.Login, uniqueCode string, wg *sync.WaitGroup) (interface{}, error) {
	defer wg.Done()
	var mtx sync.Mutex
	slog.Info(uniqueCode + " Login check auth... ")
	user, errCheck := repository.connection.CheckUsername(authData.Username)
	errors.Is(errCheck, gorm.ErrRecordNotFound)
	if errCheck != nil {
		res := helper.BuildResponse("400", errCheck.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return res, errCheck
	}

	slog.Info(uniqueCode + " Login compare password... ")
	comparePassword := ComparePassword(user.Password, []byte(authData.Password))
	if comparePassword != nil {
		res := helper.BuildResponse("400", "wrong password", helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return res, comparePassword
	}

	slog.Info(uniqueCode + " Login generating access token... ")
	mtx.Lock()
	accessToken, errAccessToken := GenerateAccessToken()
	if errAccessToken != nil {
		res := helper.BuildResponse("400", errAccessToken.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return res, errAccessToken
	}

	slog.Info(uniqueCode + " Login generating refresh token... ")
	refreshToken, errRefreshToken := GenerateRefreshToken()
	if errRefreshToken != nil {
		res := helper.BuildResponse("400", errRefreshToken.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return res, errAccessToken
	}

	updateTokenData := dto.Auth{
		RefreshToken: refreshToken,
	}
	mtx.Unlock()

	slog.Info(uniqueCode + " Login updating refresh token... ")
	updateToken := repository.connection.UpdateRefreshToken(&updateTokenData, authData.Username)
	if updateToken != nil {
		res := helper.BuildResponse("400", updateToken.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return res, updateToken
	}

	data := dto.ResponseToken{AccessToken: accessToken, RefreshToken: refreshToken}

	return data, nil
}

func (repository *authService) Register(authData dto.Register, uniqueCode string) error {
	slog.Info(uniqueCode + " Register check user username... ")
	_, errCheck := repository.connection.CheckUsername(authData.Username)
	errors.Is(errCheck, gorm.ErrDuplicatedKey)

	if errCheck == nil {
		res := helper.BuildResponse("401", "duplicate username", helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return errCheck
	}

	slog.Info(uniqueCode + " Register check user email... ")
	_, errCheckEmail := repository.connection.CheckEmail(authData.Email)
	errors.Is(errCheckEmail, gorm.ErrDuplicatedKey)

	if errCheckEmail == nil {
		res := helper.BuildResponse("401", "duplicate email", helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return errCheckEmail
	}

	slog.Info(uniqueCode + " Register hashing password... ")
	hashedPassword, errHash := HashPassword(authData.Password)
	if errHash != nil {
		res := helper.BuildResponse("400", errHash.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return errHash
	}

	dataInsert := dto.Auth{
		Username:  authData.Username,
		Email:     authData.Email,
		Password:  hashedPassword,
		Name:      authData.Name,
		StatusID:  1,
		CreatedAt: time.Now(),
	}

	slog.Info(uniqueCode + " Register insert data... ")
	insert := repository.connection.Insert(&dataInsert)
	if insert != nil {
		res := helper.BuildResponse("400", insert.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Register response ", res)
		return insert
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

func GenerateAccessToken() (string, error) {
	tokenSecret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	claims := jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resToken, err := token.SignedString([]byte(tokenSecret))

	return resToken, err
}

func GenerateRefreshToken() (string, error) {
	tokenSecret := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	claims := jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 10000).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resToken, err := token.SignedString([]byte(tokenSecret))

	return resToken, err
}

func ComparePassword(hashedPassword string, plainPassword []byte) error {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err
}
