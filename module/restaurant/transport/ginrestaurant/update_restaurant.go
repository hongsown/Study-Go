package ginrestaurant

import (
	"StudyGo/common"
	"StudyGo/component/appctx"
	restaurantbiz "StudyGo/module/restaurant/biz"
	restaurantmodel "StudyGo/module/restaurant/model"
	restaurantstorage "StudyGo/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := restaurantstorage.NewSqlStore(db)

		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}
