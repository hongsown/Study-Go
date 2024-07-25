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
			panic(common.ErrInvalidRequest(err))
		}
		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSqlStore(db)

		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}
