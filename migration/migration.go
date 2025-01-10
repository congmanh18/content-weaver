package migration

import (
	"log"

	"github.com/congmanh18/content-weaver/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Printf("Migration.....")
	err := db.AutoMigrate(
		entity.Ebook{},
		entity.EbookOutline{},
		entity.EbookContent{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
