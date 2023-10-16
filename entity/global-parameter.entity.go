package entity

type GlobalParameter struct {
	ID    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Code  string `gorm:"type:varchar(255)" json:"code"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Value string `gorm:"type:varchar(255)" json:"value"`
}
