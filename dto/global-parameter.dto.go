package dto

import "time"

type GlobalParameter struct {
	ID          int       `json:"id" form:"id"`
	Code        string    `json:"code" form:"code"`
	Name        string    `json:"name" form:"name"`
	Value       string    `json:"value" form:"value"`
	Description string    `json:"description" form:"description"`
	StatusID    int       `json:"status_id" form:"status_id"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

type GlobalParameterCreate struct {
	// Code      string    `json:"code" form:"code" validate:"required"`
	Name      string    `json:"name" form:"name" validate:"required"`
	Value     string    `json:"value" form:"value" validate:"required"`
	StatusID  int       `json:"status_id" form:"status_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type GlobalParameterAll struct {
	ID       int    `json:"id" form:"id"`
	Code     string `json:"code" form:"code" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Value    string `json:"value" form:"value" validate:"required"`
	StatusID string `json:"status" form:"status" validate:"required"`
}

type GlobalParameterUpdate struct {
	Code  string `json:"code" form:"code" validate:"required"`
	Name  string `json:"name" form:"name" validate:"required"`
	Value string `json:"value" form:"value" validate:"required"`
}
