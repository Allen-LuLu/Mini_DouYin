package dao

import (
	"Mini_DouYin/cmd/user/conf"
	"Mini_DouYin/cmd/user/model"
	"Mini_DouYin/cmd/user/pkg"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var mysqlClient *gorm.DB


// InitMysql 初始化mysql
func InitMysql(){
	cfg := conf.Cfg.MysqlCfg
	addr := fmt.Sprintf("%v:%v",cfg.Host,cfg.Port)
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",cfg.UserName,cfg.PassWord,addr,cfg.DBName)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,       // 使用单数表名
		},
	})
	if err != nil {
		panic(any(err))
	}
	mysqlClient = db
}

// CreatUser 创建用户
func CreatUser(userName,pwd string)(*model.User,error){

	pwdHash,err := pkg.HashPassword(pwd)
	if err != nil {
		return nil,err
	}

	user := &model.User{UserName: userName,PassWord: pwdHash,CreatedAT: time.Now(),DelState: 1}
	err = mysqlClient.Create(user).Error
	if err != nil {
		log.Println("create user failed")
		return nil,err
	}

	mysqlClient.Last(&user)

	return user,nil
}

// GetUserIdByUserName 通过用户名获取用户id
func GetUserIdByUserName(username string)(int64,error){
	var user model.User
	err := mysqlClient.Model(&model.User{}).Where("user_name=?",username).First(&user).Error
	if err != nil {
		return -1,err
	}
	return user.UserID,nil
}

// CheckUser 检查用户名密码是否匹配
func CheckUser(username string,password string)(int64,error){
	pwd,err := pkg.HashPassword(password)
	if err != nil {
		return -1,err
	}

	var user model.User
	err = mysqlClient.Model(&model.User{}).Where("user_name=? and password_digest=?",username,pwd).First(&user).Error
	if err != nil {
		return -1,err
	}
	return user.UserID,nil

}

// GetUserByID 通过用户ID获取用户信息
func GetUserByID(userID int64)(*model.User,error){
	var user model.User
	err := mysqlClient.Model(&model.User{}).Where("user_id=?",userID).First(&user).Error
	if err != nil {
		return nil,err
	}
	return &user,nil
}

