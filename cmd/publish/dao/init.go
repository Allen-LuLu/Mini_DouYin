package dao

// Initdb Initialize mysql and other services
func Initdb() {
	InitMysql()
	InitTencentCloud()
}
