package migration

import (
	"s-store/internal/model/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.UserEntity{})
}
