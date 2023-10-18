package entity

import "time"

type Status struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Code        string    `gorm:"type:varchar(255)" json:"code"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (Status) TableName() string {
	return "status"
}
