package database

import (
	"github.com/huihuangyu/go-rest-api-gorm-fiber/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrate the datebase and creates the comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}

	return nil
}
