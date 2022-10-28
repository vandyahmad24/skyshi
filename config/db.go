package config

import (
	"fmt"
	"log"
	"os"
	"time"
	"vandyahmad/skyshi/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var e error

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	fmt.Println("Trying to connect database :" + GetEnvVariable("MYSQL_DBNAME"))
	fmt.Println("Trying to connect MYSQL_HOST :" + GetEnvVariable("MYSQL_HOST"))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnvVariable("MYSQL_USER"),
		GetEnvVariable("MYSQL_PASSWORD"),
		GetEnvVariable("MYSQL_HOST"),
		GetEnvVariable("MYSQL_PORT"),
		GetEnvVariable("MYSQL_DBNAME"))
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if e != nil {
		panic(e)
	}
	fmt.Println("Success connect database :" + GetEnvVariable("MYSQL_DBNAME"))

	InitMigrate()
	InitSeeder()
}

func InitMigrate() {
	fmt.Println("Jalankan migration")
	DB.AutoMigrate(
		&migration.Activity{}, &migration.Todo{},
	)
	fmt.Println("Selesai migration")

}

func InitSeeder() {
	fmt.Println("Jalankan Seeder")
	// seeder := seeder.NewSeeder(DB)
	// seeder.ActivitySeeder()
	// seeder.TodoSeeder()
	fmt.Println("Selesai Seeder")
}
