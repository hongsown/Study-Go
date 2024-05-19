package restaurantstorage

import (
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil

}
