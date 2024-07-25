package restaurantbiz

import (
	"StudyGo/common"
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
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
		return common.ErrEntityNotFound(err, restaurantmodel.EntityName)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(err, restaurantmodel.EntityName)
	}
	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(err, restaurantmodel.EntityName)
	}
	return nil
}
