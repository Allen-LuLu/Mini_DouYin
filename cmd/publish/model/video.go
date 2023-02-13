package model

type Video struct {
	Id       int64  `gorm:"column:id;primaryKey,autoIncrement"`
	AuthId   int64  `gorm:"column:auth_id"`
	Title    string `gorm:"column:title"`
	PlayURL  string `gorm:"column:play_URL"`
	CoverURL string `gorm:"column:cover_URL"`
	FavCnt   int64  `gorm:"column:favorite_count"`
	ComCnt   int64  `gorm:"column:comment_count"`
	CrtTime  int64  `gorm:"column:create_time,autoCreateTime:milli"`
	DelStu   bool   `gorm:"column:del_state"`
}
