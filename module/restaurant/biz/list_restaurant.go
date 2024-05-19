package restaurantbiz

import (
	"StudyGo/common"
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
)

type ListResTauRant interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListResTauRant
}

func NewListResTauRantBiz(store ListResTauRant) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListResTauRant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
