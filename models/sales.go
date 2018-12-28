package models

import (
	"github.com/jinzhu/gorm"
)

type Sales struct {
	gorm.Model
	Note        string       `gorm:"type:varchar(255)" json:"note,omitempty"`
	ProductOuts []ProductOut `json:"productOuts,omitempty"`
}
