package restaurantbiz

import (
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
)

type CreateRestaurant interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurant
}

func NewCreateResTauRantBiz(store CreateRestaurant) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
