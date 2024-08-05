package database

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDBInstance() *gorm.DB {
	if db == nil {
		db = connectDB()
	}

	return db
}

func connectDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error()
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf(
		`user=%v password=%v host=%v port=%v database=%v sslmode=disable`,
		username, password, host, port, databaseName,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Error().Msgf("Can't connect to database: %s", err)
	}

	if err := defineEnums(db); err != nil {
		log.Error().Msgf("Can't define enums: %s", err)
	}

	return db
}

func defineEnums(db *gorm.DB) error {
	enumStatements := map[string]string{
		"seats_status":       "('available', 'booked')",
		"transaction_status": "('pending', 'ongoing', 'completed', 'cancelled')",
	}

	for typename, statement := range enumStatements {
		if !IsEnumExists(db, typename) {
			createStmt := fmt.Sprintf("CREATE TYPE %s AS ENUM %s", typename, statement)
			if err := db.Exec(createStmt).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func IsEnumExists(db *gorm.DB, enumName string) bool {
	var exists bool
	db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = ?)", enumName).Scan(&exists)

	return exists
}