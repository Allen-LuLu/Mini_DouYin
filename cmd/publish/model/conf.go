package model

type Cfg struct {
	MysqlCfg MysqlCfg `json:"mysql"`
}
type MysqlCfg struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	UsrName string `json:"username"`
	Pwd     string `json:"password"`
	DBName  string `json:"db_name"`
}
