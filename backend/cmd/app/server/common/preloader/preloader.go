package preloader

import "gorm.io/gorm"

type Preloader interface {
	Preload(tx *gorm.DB) *gorm.DB
}
