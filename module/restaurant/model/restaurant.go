package restaurantmodel

import (
	"StudyGo/common"
	"errors"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel
	Addr   string `json:"addr" gorm:"column:addr;"`
	Name   string `json:"name" gorm:"column:name;"`
	Status int    `json:"status" gorm:"column:status;default:1"` // Set default value to 1
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

// update table with field name is empty
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

var (
	ErrNameIsEmpty = errors.New("Name can not be empty")
)
