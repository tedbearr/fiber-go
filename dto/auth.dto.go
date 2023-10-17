package dto

import "time"

type Auth struct {
	ID           int       `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Email        string    `json:"email" form:"email"`
	Username     string    `json:"username" form:"username"`
	Password     string    `json:"password" form:"password"`
	RefreshToken string    `json:"refresh_token" form:"refresh_token"`
	StatusID     int       `json:"status_id" form:"status_id"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
}

type Register struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Login struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type ResponseToken struct {
	AccessToken  string `json:"accessToken" form:"accessToken"`
	RefreshToken string `json:"refreshToken" form:"refreshToken"`
}
