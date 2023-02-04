package dao

func Init()  {
	InitMysql() // 初始化mysql
	initRedis() // 初始化redis
}
