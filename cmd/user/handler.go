package main

import (
	"Mini_DouYin/cmd/user/dao"
	"Mini_DouYin/common/consts/errmsg"
	"Mini_DouYin/kitex_gen/user"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	resp = new(user.RegisterResp)
	resp.Base = new(user.BaseResp)

	_, err = dao.GetUserIdByUserName(req.Username) // 到数据库查询用户是否存在
	if !errors.Is(err, gorm.ErrRecordNotFound) { // 如果err不是ErrRecordNotFound，说明数据库中可能存在着该用户的记录，或者出现其他错误
		if err == nil { // 说明数据库中可能存在着该用户的记录，返回通知客户端修改用户名
			resp.Base.Code = consts.StatusBadRequest
			resp.Base.Errmsg = errmsg.USERALREADYEXITS
			return resp, nil
		}
		return nil, err
	}

	user, err := dao.CreatUser(req.Username, req.Password) // 数据库中创建新的用户记录

	if err != nil {
		return nil, err
	}


	resp.UserID = user.UserID
	resp.Base.Code = consts.StatusOK
	resp.Base.Errmsg = "ok"
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
// 该服务只负责加密密码并到数据库中匹配记录，匹配成功则返回用户id
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserReq) (resp *user.CheckUserResp, err error) {
	// TODO: Your code here...
	resp = new(user.CheckUserResp)

	userID, err := dao.CheckUser(req.Username, req.Password) // 到数据库中匹配记录
	if err != nil || userID < 0 {
		return resp, err
	}

	resp.UserID = userID
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
// 该服务负责查询给定的id列表相对应的用户信息
// 在api层jwt鉴权通过后才能调用该服务
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	// TODO: Your code here...
	resp = new(user.UserInfoResp)
	resp.Base = new(user.BaseResp)
	resp.Users = make([]*user.User,0)

	users, err := dao.GetUserByID(req.UserIDS)
	if err != nil {
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Errmsg = errmsg.USERNOTEXIST
		return resp, nil
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Errmsg = "ok"
	for _,u := range users{
		resp.Users = append(resp.Users,&user.User{
			Id:            u.UserID,
			Name:          u.UserName,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
		})
	}
	return
}


