package logic

import (
	"context"
	"github.com/007team/douyin-micro/user/dao/mysql"
	"github.com/007team/douyin-micro/user/dao/redis"
	"github.com/007team/douyin-micro/user/models"
	"github.com/007team/douyin-micro/user/services"
	"log"
)

// 关注列表
func (s *UserService) FollowList(ctx context.Context, request *services.FollowListRequest, response *services.FollowListResponse) error {
	userId := request.UserId
	es, err := redis.FollowList(userId)
	if err != nil {
		log.Println("FollowList failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}
	var user []models.User
	// mysql查询用户
	if len(es) != 0 {
		user, err = mysql.FollowList(es)
	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.UserList = BuildUserList(user)

	return nil

}


// 粉丝列表
func (s *UserService) FollowerList(ctx context.Context, request *services.FollowerListRequest, response *services.FollowerListResponse) error {
	userId := request.UserId
	es, err := redis.FollowerList(userId)
	if err != nil {
		log.Println("FollowList failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}
	var user []models.User
	if len(es) != 0 {
		user, err = mysql.FollowerList(es)
	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.UserList = BuildUserList(user)

	return nil
}
