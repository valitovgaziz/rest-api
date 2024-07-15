package storage

import (
	"fmt"
	"log"

	"github.com/valitovgaziz/rest-api/configs"
	"github.com/valitovgaziz/rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to the Database
func ConnectDB(configs *configs.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", configs.DBHost, configs.DBUserName, configs.DBUserPassword, configs.DBName, configs.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	fmt.Println("🚀 Connected Successfully to the Database")
}

func DropTables() {
	DB.Migrator().DropTable(
		&models.DepthOrderAsks{},
		&models.DepthOrderBids{},
		&models.OrderBook{},
		&models.HistoryOrder{},
	)
}

func Close() {
	conn, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to close the Database")
	}
	defer conn.Close()
}
