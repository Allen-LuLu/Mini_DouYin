package model

type Cfg struct{
	MysqlCfg mysqlCfg `ini:"mysql"`  // mysql相关配置
	RedisCfg redisCfg `ini:"redis"`  // redis相关配置
	UserCfg userCfg	 `ini:"UserService"` // UserService相关配置
	EtcdCfg etcdCfg  `ini:"etcd"`  // etcd相关配置
	TraceCfg traceCfg `ini:"trace"`  // trace相关配置
}

type mysqlCfg struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	UserName string `ini:"username"`
	PassWord string `ini:"password"`
	DBName string `ini:"dbName"`
}

type redisCfg struct {
	Host string `ini:"host"`
	PassWord string `ini:"password"`
	DBNum int `ini:"dbNum"`
}

type userCfg struct {
	Addr string `ini:"addr"`
	ServiceName string `ini:"serviceName"`
}

type etcdCfg struct{
	Addr string `ini:"addr"`
}

type traceCfg struct{
	ExportEndPoint string `ini:"ExportEndPoint"`
}