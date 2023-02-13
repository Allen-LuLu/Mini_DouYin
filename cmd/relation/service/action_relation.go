package service

import (
	"Mini_DouYin/cmd/relation/dao"
	"Mini_DouYin/cmd/relation/model"
	"Mini_DouYin/kitex_gen/relation"
	"context"
	"errors"
)

const (
	addRelation int32 = iota
	delRealtion
)

type ActionRelationService struct {
	ctx context.Context
}

func NewActionRelationService(ctx context.Context) *ActionRelationService {
	return &ActionRelationService{ctx: ctx}
}

func (s *ActionRelationService) ActionRelation(req *relation.ActionRelationRequest) (err error) {
	follow := &model.Follow{
		UserId:   req.UserId,
		FollowId: req.FollowId,
	}

	follower := &model.Follow{
		UserId:   req.FollowId,
		FollowId: req.UserId,
	}

	switch req.ActionType {
	case addRelation:
		exist, _ := dao.QueryRelation(s.ctx, follower)
		if exist {
			follow.IsFollow = 1
			follower.IsFollow = 1
			dao.ModifyRelation(s.ctx, follower)
		}
		err = dao.AddRelation(s.ctx, follow)
	case delRealtion:
		exist, _ := dao.QueryRelation(s.ctx, follower)
		if exist {
			follower.IsFollow = 0
			dao.ModifyRelation(s.ctx, follower)
		}
		err = dao.DelRelation(s.ctx, follow)
	default:
		err = errors.New("不支持的关系操作")
	}
	return
}
