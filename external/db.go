package external

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sriharivishnu/shopify-challenge/config"
	models "github.com/sriharivishnu/shopify-challenge/models/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func createDatabase() error {
	// host := config.Config.DATABASE_HOST
	// port := config.Config.DATABASE_PORT
	// user := config.Config.DATABASE_USER
	// pass := config.Config.DATABASE_PASSWORD
	database := config.Config.DATABASE_NAME

	log.Println("Creating Database...")
	var err error
	DbConn, err = gorm.Open(sqlite.Open(database), &gorm.Config{})
	return err
}

func createSchema() {
	log.Println("Creating Schema...")
	DbConn.AutoMigrate(&models.Item{})
	DbConn.AutoMigrate(&models.Warehouse{})
	log.Println("Done creating schema")
}

func Init() {
	err := createDatabase()
	retryCount := 10
	for err != nil && retryCount >= 0 {
		log.Printf("Attempted to connect to database and failed: %v retryCount: %d", err, retryCount)
		retryCount--
		time.Sleep(time.Duration(11-retryCount) * time.Second)
		err = createDatabase()
	}
	if err != nil {
		panic("Could not connect to database!")
	}
	log.Println("Connected to DB host")

	log.Println("Verifying schema...")
	createSchema()
	log.Println("Database is connected and ready!")
}
