package model


import "time"

type User struct{
	UserID int64 `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	PassWord string `gorm:"column:password_digest"`
	FollowCount int64 `gorm:"column:follow_count"`
	FollowerCount int64 `gorm:"column:follower_count"`
	DelState int `gorm:"column:del_state"`
	CreatedAT time.Time `gorm:"column:create_time"`
}