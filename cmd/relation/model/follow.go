package model

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserId    int64 `gorm:"colomn:user_id"`
	FollowId  int64 `gorm:"colomn:follow_id"`
	IsFollow  int8  `gorm:"colomn:is_follow"`
	deletedAt gorm.DeletedAt
}

func (f *Follow) TableName() string {
	return "follow"
}
