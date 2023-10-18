package entity

import "time"

type User struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name         string `gorm:"type:varchar(255)" json:"name"`
	Username     string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Email        string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password     string `gorm:"type:varchar(255)" json:"password"`
	RefreshToken string `gorm:"type:varchar(255)" json:"refresh_token"`
	StatusID     uint64 `gorm:"not null;default:1" json:"-"`
	// Status       Status    `gorm:"foreignkey:StatusID;constraint:onDelete:CASCADE" json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
