package database

import (
	"log"
	"os"

	"github.com/CodeATM/fibre-gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// Import the modernc SQLite driver so that it registers itself.
	_ "modernc.org/sqlite"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	// DSN for your database. You can also add extra query parameters if needed.
	dsn := "api.db"

	// Use sqlite.New with a configuration that specifies the modernc driver.
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DSN:        dsn,
		DriverName: "sqlite", // Use "sqlite" (modernc's driver) instead of the default "sqlite3"
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to db: ", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to db")

	log.Println("Running migration")
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}); err != nil {
		log.Fatal("Failed to migrate: ", err.Error())
		os.Exit(3)
	}

	Database = DbInstance{Db: db}
}
