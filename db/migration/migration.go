package migration

import (
	"log"
	"sync"

	"gorm.io/gorm"
)

var onlyOnce sync.Once

func Migrate(connection *gorm.DB) {

	onlyOnce.Do(func() {

		log.Println("Migrating the database...")

		connection.AutoMigrate()

	})

}
