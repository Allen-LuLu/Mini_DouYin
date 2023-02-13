package rpc

import (
	"Mini_DouYin/cmd/api/biz/conf"
	"Mini_DouYin/common/mw"
	"Mini_DouYin/kitex_gen/user"
	"Mini_DouYin/kitex_gen/user/userservice"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var userClient userservice.Client

func initUser(){
	r,err := etcd.NewEtcdResolver([]string{conf.Cfg.EtcdCfg.Addr})
	if err != nil {
		panic(any(err))
	}


	c,err := userservice.NewClient(
		conf.Cfg.UserCfg.ServiceName,
		client.WithResolver(r),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.Cfg.ApiCfg.ServiceName}),
	)
	if err != nil {
		panic(any(err))
	}
	userClient = c
}

func Register(ctx context.Context,req *user.RegisterReq)(*user.RegisterResp,error){
	resp,err := userClient.Register(ctx,req)
	if err != nil {
		log.Println("rpc register failed")
		return nil,err
	}
	return resp,nil
}

func CheckUser(ctx context.Context,req *user.CheckUserReq)(*user.CheckUserResp,error){
	resp,err := userClient.CheckUser(ctx,req)
	if err != nil {
		log.Println("rpc login failed")
		return nil,err
	}
	return resp,nil
}

func GetUserInfo(ctx context.Context,req *user.UserInfoReq)(*user.UserInfoResp,error){
	resp,err := userClient.UserInfo(ctx,req)
	if err != nil {
		log.Println("rpc login failed")
		return nil,err
	}
	return resp,nil
}