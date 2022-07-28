package models

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model
	ID uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
}
