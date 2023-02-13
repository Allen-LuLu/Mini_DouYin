package dao

import (
	"Mini_DouYin/common/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.RelationMySQLDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

}

func init() {
	initDB()
}
