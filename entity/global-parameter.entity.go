package entity

import "time"

type GlobalParameter struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Code        string `gorm:"uniqueIndex;type:varchar(255)" json:"code"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Value       string `gorm:"type:varchar(255)" json:"value"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	StatusID    uint64 `gorm:"not null;default:1" json:"-"`
	// Status    Status    `gorm:"foreignkey:StatusID;constraint:onDelete:CASCADE" json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (GlobalParameter) TableName() string {
	return "global_parameter"
}
