// Code generated by hertz generator.

package main

import (
	"Mini_DouYin/cmd/api/biz/conf"
	"Mini_DouYin/cmd/api/biz/dao"
	"Mini_DouYin/cmd/api/biz/mw"
	"Mini_DouYin/cmd/api/biz/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
)

func main() {
	conf.Init() // 配置初始化
	mw.InitJWT() // jwt插件初始化
	dao.Init() // 数据库初始化
	rpc.Init() // rpc初始化

	tracer,cfg := tracing.NewServerTracer()


	h := server.New(
		server.WithHostPorts(conf.Cfg.ApiCfg.Addr),
		server.WithHandleMethodNotAllowed(true),
		tracer,
	)


	pprof.Register(h)
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
