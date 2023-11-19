package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var SqlDb *gorm.DB

const (
	sqlAddr         = "127.0.0.1:3306"
	sqlUser         = "root"
	sqlPswd         = ""
	sqlDataBaseName = "extreme"
)

func InitDb() {
	initMySQL()
}

func initMySQL() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb3&parseTime=True&loc=Local", sqlUser, sqlPswd, sqlAddr, sqlDataBaseName)
	log.Printf("Connecting to MySQL server: dsn = %v\n", dsn)

	var err error
	SqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error occoured when connecting to the data base! error: %e", err)
	}

	err = SqlDb.AutoMigrate(&User{}, &RateRecord{}, &CalcHistory{})
	if err != nil {
		log.Fatalf("failed to initialize database: %e", err)
	}

	// Print a success message
	fmt.Println("Database initialized successfully")
}
