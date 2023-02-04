package model


type Cfg struct{
	UserCfg userCfg	 `ini:"UserService"` // UserService相关配置
	EtcdCfg etcdCfg  `ini:"etcd"` // etcd相关配置
	ApiCfg apiCfg `ini:"api"` // api相关配置
}


type userCfg struct {
	Addr string `ini:"addr"`
	ServiceName string `ini:"serviceName"`
}

type etcdCfg struct{
	Addr string `ini:"addr"`
}


type apiCfg struct {
	Addr string `ini:"addr"`
	ServiceName string `ini:"serviceName"`
}
