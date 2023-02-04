package main

import (
	"Mini_DouYin/cmd/user/conf"
	"Mini_DouYin/cmd/user/dao"
	"Mini_DouYin/common/mw"
	user "Mini_DouYin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	conf.Init()  //配置初始化
	dao.Init() //数据库初始化

	r,err := etcd.NewEtcdRegistry([]string{conf.Cfg.EtcdCfg.Addr})  //服务注册
	if err != nil {
		panic(any(err))
	}

	addr,err := net.ResolveTCPAddr("tcp",conf.Cfg.UserCfg.Addr)  //绑定服务地址
	if err != nil {
		panic(any(err))
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections:1000,MaxQPS: 100}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.Cfg.UserCfg.ServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
