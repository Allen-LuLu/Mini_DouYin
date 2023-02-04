package dao

import (
	"Mini_DouYin/cmd/user/conf"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

var redisClient *redis.Client

// initRedis 初始化redis
func initRedis(){
	cfg := conf.Cfg.RedisCfg
	db := redis.NewClient(&redis.Options{
		Addr: cfg.Host,
		Password: cfg.PassWord,
		DB: cfg.DBNum,
	})
	_,err := db.Ping().Result()
	if err != nil{
		panic(any(err))
	}
	redisClient = db
}

// StoreToken redis中以string格式存储token
func StoreToken(userID int64,tokenString string,expire time.Duration)error{
	return redisClient.SetNX(strconv.Itoa(int(userID)),tokenString,expire).Err()
}

// QueryToken 查询token
func QueryToken(userID int64)(token string,err error){
	return redisClient.Get(strconv.Itoa(int(userID))).Result()
}
