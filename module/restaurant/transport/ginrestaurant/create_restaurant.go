package ginrestaurant

import (
	"StudyGo/common"
	"StudyGo/component/appctx"
	restaurantbiz "StudyGo/module/restaurant/biz"
	restaurantmodel "StudyGo/module/restaurant/model"
	restaurantstorage "StudyGo/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := restaurantstorage.NewSqlStore(db)

		biz := restaurantbiz.NewCreateResTauRantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRes(data.Id))
	}
}
