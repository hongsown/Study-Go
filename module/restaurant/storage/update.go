package restaurantstorage

import (
	restaurantmodel "StudyGo/module/restaurant/model"
	"context"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ? AND status = 1", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
