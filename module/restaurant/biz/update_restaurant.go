package restaurantbiz

import (
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}
type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	dataExist, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if dataExist.Status == 0 {
		return errors.New("data dont exist")
	}
	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}
	return nil
}
