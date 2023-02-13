package service

import (
	"Mini_DouYin/cmd/relation/dao"
	"Mini_DouYin/kitex_gen/relation"
	"context"
	"errors"
)

const (
	listAllFollow int32 = iota
	listAllFollower
	listAllFriend
)

type ListAllOperateService struct {
	ctx context.Context
}

func NewListAllOperateService(ctx context.Context) *ListAllOperateService {
	return &ListAllOperateService{ctx: ctx}
}

func (s *ListAllOperateService) ListAllOperate(req *relation.ListAllOperateRequest) (IDs []int64, err error) {
	switch req.OperateType {
	case listAllFollow:
		IDs, err = dao.QueryUserAllFollows(s.ctx, req.UserId)
	case listAllFollower:
		IDs, err = dao.QueryUserAllFollowers(s.ctx, req.UserId)
	case listAllFriend:
		IDs, err = dao.QueryUserAllFriends(s.ctx, req.UserId)
	default:
		err = errors.New("不支持这种查询操作")
	}
	return
}
