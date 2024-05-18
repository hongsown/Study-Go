package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"` //tag
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

// update table with field name is empty
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {

	dsn := os.Getenv("MYSQL_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)

	//GET
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//POST
	v1 := r.Group("/v1")
	restaurants := v1.Group("restaurants")
	restaurants.POST("/", func(c *gin.Context) {
		var data Restaurant
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	//GET By ID
	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data Restaurant
		db.Where("id = ?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	//GET ALL
	restaurants.GET("", func(c *gin.Context) {
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var pagingData Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		if pagingData.Limit <= 0 {
			pagingData.Limit = 10
		}
		var data []Restaurant
		db.Offset((pagingData.Page - 1) * pagingData.Limit).
			Order("id desc").
			Limit(pagingData.Limit).
			Find(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	//UPDATE
	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Where("id = ?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	//DELETE
	restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)
		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	})
	r.Run()
	//CREATE
	//newRestaurant := Restaurant{Name: "Banh Beo", Addr: "123 abc street"}
	//
	//db.Create(&newRestaurant)

	//READ
	//var myRestaurant Restaurant
	//if err := db.Where("id = ?", 1).First(&myRestaurant).Error; err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(myRestaurant)

	//UPDATE
	//newDataUpdate := "Banh Canh"
	//updatedData := RestaurantUpdate{Name: &newDataUpdate}
	//if err := db.Where("id = ?", 3).Updates(&updatedData).Error; err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(myRestaurant)

	//DELETE
	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(myRestaurant)
}
