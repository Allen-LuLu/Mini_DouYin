package model

type Cfg struct {
	MysqlCfg   MysqlCfg   `json:"mysql"`
	TencentCfg TencentCfg `json:"tencentCloud"`
}
type MysqlCfg struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	UsrName string `json:"username"`
	Pwd     string `json:"password"`
	DBName  string `json:"db_name"`
}
type TencentCfg struct {
	BucketName string `json:"bucketName"`
	BucketID   string `json:"bucketID"`
	CosRegion  string `json:"cosRegion"`
	SecretId   string `json:"secretId"`
	SecretKey  string `json:"secretKey"`
}
