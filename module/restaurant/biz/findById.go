package restaurantbiz

import (
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
)

type FindByIdStore interface {
	FindDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}
type findRestaurantBiz struct {
	store FindByIdStore
}

func NewFindRestaurantBiz(store FindByIdStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}

func (biz *findRestaurantBiz) FindByIdStore(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {

	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}
