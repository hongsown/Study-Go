package ginrestaurant

import (
	"StudyGo/component/appctx"
	restaurantbiz "StudyGo/module/restaurant/biz"
	restaurantstorage "StudyGo/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDataById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewFindRestaurantBiz(store)
		result, err := biz.FindByIdStore(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
