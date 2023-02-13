package dao

import (
	"Mini_DouYin/cmd/relation/model"
	"context"
)

func AddRelation(ctx context.Context, relation *model.Follow) error {
	if err := DB.WithContext(ctx).Create(relation).Error; err != nil {
		return err
	}
	return nil
}

func AddBatchRelations(ctx context.Context, relations []*model.Follow) error {
	if err := DB.WithContext(ctx).Create(relations).Error; err != nil {
		return err
	}
	return nil
}

func DelRelation(ctx context.Context, relation *model.Follow) error {
	if err := DB.WithContext(ctx).Where("user_id = ? && follow_id = ?", relation.UserId, relation.FollowId).Delete(&model.Follow{}).Error; err != nil {
		return err
	}
	return nil
}

func QueryRelation(ctx context.Context, relation *model.Follow) (bool, error) {
	var total int64
	result := DB.WithContext(ctx).Where("user_id = ? && follow_id = ?", relation.UserId, relation.FollowId).Count(&total)
	err := result.Error
	if err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func ModifyRelation(ctx context.Context, relation *model.Follow) error {
	err := DB.WithContext(ctx).Where("user_id = ? && follow_id = ?", relation.UserId, relation.FollowId).Updates(map[string]interface{}{"is_follow": relation.IsFollow}).Error
	return err
}

func QueryUserAllFollows(ctx context.Context, userId int64) (followsIDs []int64, err error) {
	err = DB.WithContext(ctx).Model(&model.Follow{}).Select("follow_id").Where("user_id = ?", userId).Find(followsIDs).Error
	return
}

func QueryUserAllFollowers(ctx context.Context, userId int64) (followersIDs []int64, err error) {
	err = DB.WithContext(ctx).Model(&model.Follow{}).Select("user_id").Where("follow_id = ?", userId).Find(followersIDs).Error
	return
}

func QueryUserAllFriends(ctx context.Context, userId int64) (friendsIDs []int64, err error) {
	err = DB.WithContext(ctx).Model(&model.Follow{}).Select("follow_id").Where("user_id = ? && is_follow = ?", userId, 1).Find(friendsIDs).Error
	return
}
