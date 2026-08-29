package conn

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectMySQL(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err == nil {
			log.Println("✅ Connected to MySQL successfully")
			return db, nil
		}

		log.Printf("⏳ Waiting for MySQL... attempt %d\n", i)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to MySQL after retries: %v", err)
}
