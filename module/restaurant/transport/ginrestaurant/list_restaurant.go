package ginrestaurant

import (
	"StudyGo/common"
	"StudyGo/component/appctx"
	restaurantbiz "StudyGo/module/restaurant/biz"
	restaurantmodel "StudyGo/module/restaurant/model"
	restaurantstorage "StudyGo/module/restaurant/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		pagingData.Fulfill()
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewListResTauRantBiz(store)
		result, err := biz.ListResTauRant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		fmt.Println("result: ", result)

		for i := range result {
			result[i].Mask(false)

		}

		c.JSON(http.StatusOK, common.NewSuccessRes(result, pagingData, filter))
	}
}
