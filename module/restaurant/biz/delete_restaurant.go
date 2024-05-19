package restaurantbiz

import (
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
	"errors"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}
type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {

	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}
	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
