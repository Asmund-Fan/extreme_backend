package controllers

import (
	"calculator_backend/data"
	"calculator_backend/db"
	"calculator_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

func Register(c *gin.Context) {
	uName := c.Query("user_name")
	pswd := c.Query("pswd")

	if data.CheckIfNameExist(uName) {
		c.JSON(501, "")
		c.Abort()
		return
	}

	newUsr := db.User{
		UserName: uName,
		Password: pswd,
	}

	user := data.AddNewUser(newUsr)
	if user == nil {
		c.JSON(502, "")
	}
	c.JSON(200, models.UserResponse{Data: newUsr})
}

func Login(c *gin.Context) {
	uName := c.Query("user_name")
	pswd := c.Query("pswd")

	isOk, user := data.CheckLogin(uName, pswd)

	if isOk {
		c.JSON(200, models.UserResponse{Data: *user})
	} else {
		c.JSON(501, "")
	}
}

func GetHistories(c *gin.Context) {
	uid := c.Query("user_id")
	searchHistories := data.SearchHistoriesById(uid)

	c.JSON(200, models.CalcHistoriesResponse{Data: searchHistories})
}

func AddHistory(c *gin.Context) {
	uid := c.Query("user_id")
	expression := c.Query("expression")
	result := c.Query("result")

	data.AddSearchHistories(db.CalcHistory{
		HistId:     uuid.New().String(),
		Time:       time.Now().Unix(),
		Uuid:       uid,
		Expression: expression,
		Result:     result,
	})

	c.JSON(200, "")
}

func GetRateRecord(c *gin.Context) {
	// Get the uuid from the query string
	uid := c.Query("user_id")
	// Search the rate record by the uuid
	searchRateRecord, isOK := data.SearchRateRecordById(uid)
	// Return the JSON response with the status code 200
	if isOK {
		c.JSON(200, models.RateRecordResponse{Data: searchRateRecord})
	} else {
		c.JSON(400, "")
	}
}

func AddRateRecord(c *gin.Context) {
	// Get the uuid and the rates from the query string
	var rateRecord db.RateRecord
	log.Printf("%v", c.Request)
	if err := c.ShouldBindJSON(&rateRecord); err != nil {
		c.JSON(400, "") // 请求格式不正确
		return
	}
	// 将 rateRecord 对象添加到数据库中
	if data.AddSearchRateRecords(rateRecord) {
		c.JSON(200, "") // 请求成功
	} else {
		c.JSON(502, "") // 数据库操作失败
	}
}
