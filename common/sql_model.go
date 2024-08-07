package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"` //tag
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
