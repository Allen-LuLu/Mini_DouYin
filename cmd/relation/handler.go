package main

import (
	"Mini_DouYin/cmd/relation/service"
	"Mini_DouYin/kitex_gen/relation"
	"context"
	"errors"
	"time"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

const (
	SuccessCode int64 = 0
	FailCode    int64 = 1
)

func BuildBaseResp(err error) *relation.BaseResp {
	if err == nil {
		return &relation.BaseResp{
			StatusCode:  SuccessCode,
			ServiceTime: time.Now().Unix(),
		}
	}
	return &relation.BaseResp{
		StatusCode:    FailCode,
		StatusMessage: err.Error(),
		ServiceTime:   time.Now().Unix(),
	}
}

// ActionRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ActionRelation(ctx context.Context, req *relation.ActionRelationRequest) (resp *relation.ActionRelationResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.ActionRelationResponse)

	if req.UserId < 0 || req.FollowId < 0 {
		resp.BaseResp = BuildBaseResp(errors.New("用户id非法"))
		return
	}

	err = service.NewActionRelationService(ctx).ActionRelation(req)
	resp.BaseResp = BuildBaseResp(err)

	return
}

// ListAllOperate implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListAllOperate(ctx context.Context, req *relation.ListAllOperateRequest) (resp *relation.ListAllOperateResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.ListAllOperateResponse)

	if req.UserId < 0 {
		resp.BaseResp = BuildBaseResp(errors.New("用户id非法"))
		return
	}

	resp.UserIds, err = service.NewListAllOperateService(ctx).ListAllOperate(req)
	resp.BaseResp = BuildBaseResp(err)

	return
}
