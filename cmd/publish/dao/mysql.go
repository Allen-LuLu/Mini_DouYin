package dao

import (
	"Mini_DouYin/cmd/publish/conf"
	"Mini_DouYin/cmd/publish/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var mysqlClient *gorm.DB

// InitMysql mysql initialization and connect by gorm
func InitMysql() {
	mysqlCfg := conf.Cfg.MysqlCfg
	addr := fmt.Sprintf("%v:%v", mysqlCfg.Host, mysqlCfg.Port)
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", mysqlCfg.UsrName, mysqlCfg.Pwd, addr, mysqlCfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(any(err))
	}
	mysqlClient = db
}

// AddVideo Put video info to mysql
func AddVideo(authId int64, title string, playURL string, coverURL string) (*model.Video, error) {

	//Initialize
	video := &model.Video{AuthId: authId, Title: title, PlayURL: playURL, CoverURL: coverURL}

	//Generate the rest data

	video.FavCnt = 0
	video.ComCnt = 0
	video.DelStu = false
	log.Println("Video data generated\n ", video)

	//Create row
	err := mysqlClient.Create(video).Error
	if err != nil {
		log.Printf("Save video %v failed\n", playURL)
		return nil, err
	}

	mysqlClient.Last(&video)

	return video, nil
}

// QueryVideoByAuthID Query all videos by this author and return the slices of video struct
func QueryVideoByAuthID(userId int64) ([]*model.Video, error) {
	var videos []*model.Video
	res := mysqlClient.Where(&model.Video{AuthId: userId}).Find(&videos)
	//SELECT * FROM VIDEO WHERE AUTHID=USERID
	if res.Error != nil {
		log.Println(res.Error)
		return nil, res.Error
	}
	return videos, nil
}
