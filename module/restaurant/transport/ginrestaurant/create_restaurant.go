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
		//go func() {
		//	defer common.AppRecover()
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSqlStore(db)

		biz := restaurantbiz.NewCreateResTauRantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data.FakeId.String()))
	}
}
