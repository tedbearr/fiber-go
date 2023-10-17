package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/tedbearr/go-learn/helper"
)

func Jwt(ctx *fiber.Ctx) error {
	secretKey := []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
	getHeader := ctx.Request().Header.Peek("Authorization")
	authHeader := string(getHeader)

	if authHeader == "" {
		res := helper.BuildResponse("400", "token not found", helper.EmptyObj{})
		return ctx.JSON(res)
	}

	if !strings.Contains(authHeader, "Bearer") {
		res := helper.BuildResponse("400", "invalid token", helper.EmptyObj{})
		return ctx.JSON(res)
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

	token, errParse := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if errParse != nil {
		res := helper.BuildResponse("401", errParse.Error(), helper.EmptyObj{})
		return ctx.JSON(res)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if expirationTime.Before(time.Now()) {
			res := helper.BuildResponse("400", "token is expired", helper.EmptyObj{})
			return ctx.JSON(res)
		}
	} else {
		res := helper.BuildResponse("400", "token is not valid", helper.EmptyObj{})
		return ctx.JSON(res)
	}

	return ctx.Next()
}
