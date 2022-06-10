package logic

import (
	"context"
	"github.com/007team/douyin-micro/user/pkg/jwt"
	"log"

	"github.com/007team/douyin-micro/user/dao/mysql"
	"github.com/007team/douyin-micro/user/dao/redis"
	"github.com/007team/douyin-micro/user/models"
	"github.com/007team/douyin-micro/user/services"
)

// 关注操作
func (s *UserService) RelationAction(ctx context.Context, request *services.RelationActionRequest, response *services.RelationActionResponse) error {
	userId := request.UserId
	toUserId := request.ToUserId
	actionType := request.ActionType

	if actionType == 1 {
		// 关注操作
		err := redis.FollowAction(userId, toUserId)
		if err != nil {
			log.Println("logic.RelationAction failed", err)
			response.StatusCode = 1
			response.StatusMsg = "关注失败"
			return nil

		}
		// 在对方的粉丝列表里添加我
		err = redis.FollowerActionToUser(toUserId, userId)
		if err != nil {
			log.Println("logic.RelationAction failed", err)
			response.StatusCode = 1
			response.StatusMsg = "关注失败"
			return nil
		}
		response.StatusCode = 0
		response.StatusMsg = "关注成功"
		return nil

	} else if actionType == 2 {
		// 取消关注
		// 在我的关注列表中删除对方用户
		err := redis.UnFollowAction(userId, toUserId)
		if err != nil {
			log.Println("logic.RelationAction failed", err)
			response.StatusCode = 1
			response.StatusMsg = "关注失败"
			return nil
		}
		// 在对方的粉丝列表中删除我的id
		err = redis.UnFollowerActionToUser(userId, toUserId)
		if err != nil {
			log.Println("logic.RelationAction failed", err)
			response.StatusCode = 1
			response.StatusMsg = "关注失败"
			return nil
		}
		response.StatusCode = 0
		response.StatusMsg = "取关成功"
		return nil
	}

	return nil
}

// 关注列表
func (s *UserService) FollowList(ctx context.Context, request *services.FollowListRequest, response *services.FollowListResponse) error {
	m, _ := jwt.ParseToken(request.Token)
	// 我的id
	myId := m.UserID
	log.Println(myId)

	userId := request.UserId
	es, err := redis.FollowList(userId)
	if err != nil {
		log.Println("FollowList failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}
	var userlist []*models.User
	// mysql查询用户
	if len(es) != 0 {
		userlist, err = mysql.FollowListPointer(es)
	}

	// 从redis中读取其他字段
	for i, user := range userlist {
		userlist[i].FollowCount, err = redis.UserFollowCount(user.Id)
		if err != nil {
			log.Println("redis.UserFollowCount(user.Id) failed", err)
			response.StatusCode = 1
			response.StatusMsg = "服务器繁忙，请稍后再试"
			return nil
		}
		userlist[i].FollowerCount, err = redis.UserFollowerCount(user.Id)
		if err != nil {
			log.Println("redis.UserFollowerCount(user.Id) failed", err)
			response.StatusCode = 1
			response.StatusMsg = "服务器繁忙，请稍后再试"
			return nil
		}

		userlist[i].IsFollow = true

		// log.Println(user)

	}
	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.UserList = BuildUserList(userlist)

	return nil

}

// 粉丝列表
func (s *UserService) FollowerList(ctx context.Context, request *services.FollowerListRequest, response *services.FollowerListResponse) error {
	m, _ := jwt.ParseToken(request.Token)
	// 我的id
	myId := m.UserID

	userId := request.UserId
	es, err := redis.FollowerList(userId)
	if err != nil {
		log.Println("FollowList failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}
	var userlist []*models.User
	if len(es) != 0 {
		userlist, err = mysql.FollowerListPointer(es)
	}

	// 从redis中读取其他字段
	for i, user := range userlist {
		userlist[i].FollowCount, err = redis.UserFollowCount(user.Id)
		if err != nil {
			log.Println("redis.UserFollowCount(user.Id) failed", err)
			response.StatusCode = 1
			response.StatusMsg = "服务器繁忙，请稍后再试"
			return nil
		}
		userlist[i].FollowerCount, err = redis.UserFollowerCount(user.Id)
		if err != nil {
			log.Println("redis.UserFollowerCount(user.Id) failed", err)
			response.StatusCode = 1
			response.StatusMsg = "服务器繁忙，请稍后再试"
			return nil
		}

		// “我”是否关注了这个用户
		userlist[i].IsFollow, err = redis.IsFollowUser(user, myId)
		if err != nil {
			log.Println("redis.IsFollowUser(user, myUserId) failed", err)
			response.StatusCode = 1
			response.StatusMsg = "服务器繁忙，请稍后再试"
			return nil
		}

	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.UserList = BuildUserList(userlist)

	return nil
}
