package main

import (
	"StudyGo/component/appctx"
	"StudyGo/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {

	dsn := os.Getenv("MYSQL_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()
	//GET
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	appCtx := appctx.NewAppContext(db)

	//POST
	v1 := r.Group("/v1")
	restaurants := v1.Group("restaurants")
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appCtx))
	//GET By ID
	restaurants.GET("/:id", ginrestaurant.GetDataById(appCtx))
	////GET ALL
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	////UPDATE
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
	////DELETE
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
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
