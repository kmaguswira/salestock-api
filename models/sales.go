package models

import (
	"github.com/jinzhu/gorm"
)

type Sales struct {
	gorm.Model
	ProductOuts []ProductOut `json:"productOuts,omitempty"`
}
