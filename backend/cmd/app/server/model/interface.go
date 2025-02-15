package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Baser interface {
	schema.Tabler
	Alive(tx *gorm.DB) *gorm.DB
}
