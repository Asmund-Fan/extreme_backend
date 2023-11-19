package data

import (
	"calculator_backend/db"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func CheckIfNameExist(uName string) bool {
	err := db.SqlDb.Where("user_name = ?", uName).First(db.User{}).Error
	if fmt.Sprintf("%v", err) != "record not found" && err == nil {
		return true
	}
	return false
}

func AddNewUser(user db.User) *db.User {
	user.UUID = uuid.New().String()

	err := db.SqlDb.Create(&user).Error

	if err != nil {
		log.Fatalf("failed to insert user: %e", err)
		return nil
	}

	return &user
}

func CheckLogin(usrName string, pswd string) (bool, *db.User) {
	var dbUsr db.User
	err := db.SqlDb.Where("user_name = ?", usrName).First(&dbUsr).Error
	if err != nil || dbUsr.Password != pswd {
		return false, nil
	}
	return true, &dbUsr
}

func SearchHistoriesById(uid string) []db.CalcHistory {
	var histories []db.CalcHistory                                                  // Declare a slice of CalcHistories to store the query results
	err := db.SqlDb.Where("uuid = ?", uid).Order("time asc").Find(&histories).Error // Use the Where and Find methods to perform the query and assign the error to err

	if err != nil {
		log.Fatalf("SearchHistoriesById: failed to query database: %e", err)
		return []db.CalcHistory{}
	}

	return histories
}

func AddSearchHistories(record db.CalcHistory) {
	// Insert the record into the calc_histories table
	err := db.SqlDb.Create(&record).Error

	if err != nil {
		log.Fatalf("failed to insert history: %e", err)
	}

	// Count the number of records with the same uuid
	var count int64
	db.SqlDb.Model(&db.CalcHistory{}).Where("uuid = ?", record.Uuid).Count(&count)

	// If the count is greater than 10, delete the oldest record
	if count > 10 {
		var oldest db.CalcHistory
		db.SqlDb.Where("uuid = ?", record.Uuid).Order("time asc").First(&oldest)
		db.SqlDb.Delete(&oldest)
	}
}

func SearchRateRecordById(uid string) (db.RateRecord, bool) {
	var rateRecord db.RateRecord
	err := db.SqlDb.Where("uuid = ?", uid).First(&rateRecord).Error
	// If there is an error, log it and return an empty RateRecord
	if err != nil {
		log.Printf("SearchRateRecordById: failed to query database: %e", err)
		return db.RateRecord{}, false
	}

	// Return the query result
	return rateRecord, true
}

func AddSearchRateRecords(record db.RateRecord) bool {
	// Insert the record into the rate_records table
	err := db.SqlDb.Save(&record).Error

	if err != nil {
		log.Printf("failed to insert rate record: %e", err)
		return false
	}
	return true
}
