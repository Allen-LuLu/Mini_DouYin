package main

import (
	"Mini_DouYin/cmd/user/dao"
	"Mini_DouYin/common/consts/errmsg"
	token2 "Mini_DouYin/common/pkg/token"
	"Mini_DouYin/kitex_gen/user"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	resp = new(user.RegisterResp)
	resp.Base = new(user.BaseResp)

	_,err = dao.GetUserIdByUserName(req.Username)  // 查询用户名是否已存在数据库
	if !errors.Is(err,gorm.ErrRecordNotFound){  // 如果err不是找不到，说明数据库中可能已经存在相同的用户名
		if err == nil { // err为nil，说明数据库中可能已经存在相同的用户名
			resp.Base.Code = consts.StatusBadRequest
			resp.Base.Errmsg = errmsg.USERALREADYEXITS
			return resp,nil
		}
		return nil,err
	}

	user,err := dao.CreatUser(req.Username,req.Password)  // 创建用户

	if err != nil {
		return nil,err
	}

	token,err := token2.GenToken(time.Now(),user.UserID,nil) // 生成token
	if err != nil {
		resp.Base.Code = consts.StatusInternalServerError
		resp.Base.Errmsg = errmsg.GENTOKENFAILED
		return resp,nil
	}

	err = dao.StoreToken(user.UserID,token,token2.Expire*time.Second) // 保存token
	if err != nil {
		resp.Base.Code = consts.StatusInternalServerError
		resp.Base.Errmsg = errmsg.STORETOKENFAILED
		return resp,nil
	}

	resp.Base.Code = consts.StatusOK
	resp.Base.Errmsg = "ok"
	resp.UserID = user.UserID
	resp.Token = token
	return resp,nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	resp = new(user.LoginResp)
	resp.Base = new(user.BaseResp)

	userID,err := dao.CheckUser(req.Username,req.Password) //检查用户名密码是否匹配
	if err != nil || userID < 0{
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Errmsg = errmsg.CHECKUSERFAILED
		return resp,nil
	}

	token,err := dao.QueryToken(userID) // 查询token是否过期
	if err != nil {
		if !errors.Is(err,redis.Nil){ // 如果err是redis.Nil，说明redis中找不到该用户的token了，需要重新生成，其它错误则直接返回
			resp.Base.Code = consts.StatusInternalServerError
			resp.Base.Errmsg = errmsg.QUERYTOKEN
			return resp,nil
		}
	}

	if token == ""{
		token,_ := token2.GenToken(time.Now(),userID,nil) // 重新生成token
		err = dao.StoreToken(userID,token,token2.Expire*time.Second) // 存储token
		if err != nil {
			resp.Base.Code = consts.StatusInternalServerError
			resp.Base.Errmsg = errmsg.STORETOKENFAILED
			return resp,nil
		}
	}

	resp.Base.Code = consts.StatusOK
	resp.Base.Errmsg = "ok"
	resp.UserID = userID
	resp.Token = token

	return resp,nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	// TODO: Your code here...
	resp = new(user.UserInfoResp)
	resp.Base = new(user.BaseResp)
	resp.User = new(user.User)

	token,err := dao.QueryToken(req.UserID)  // 根据用户id查询token
	if err != nil {
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Errmsg = errmsg.AUTHFAILED
		return resp,nil
	}

	if token != req.Token{ // 校验客户端传来的token
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Errmsg = errmsg.AUTHFAILED
		return resp,nil
	}

	claim,_ := token2.ParseToken(token) // 解析token

	user,err := dao.GetUserByID(claim.UserId) // 获取用户信息
	if err != nil {
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Errmsg = errmsg.USERNOTEXIST
		return resp,nil
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Errmsg = "ok"
	resp.User.Id = user.UserID
	resp.User.Name = user.UserName
	resp.User.FollowCount = user.FollowCount
	resp.User.FollowerCount = user.FollowerCount
	return resp,nil
}
