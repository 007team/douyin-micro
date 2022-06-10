package logic

import (
	"context"
	"github.com/007team/douyin-micro/comment/dao/mysql"
	"github.com/007team/douyin-micro/comment/dao/redis"
	"github.com/007team/douyin-micro/comment/models"
	"github.com/007team/douyin-micro/comment/pkg/jwt"
	"github.com/007team/douyin-micro/comment/services"
	"log"
)

func BuildUser(item models.User) *services.User {
	userModel := services.User{
		Id:            item.Id,
		Name:          item.Name,
		FollowCount:   item.FollowCount,
		FollowerCount: item.FollowerCount,
		IsFollow:      item.IsFollow,
	}
	return &userModel
}

func BuildComment(item models.Comment) *services.Comment {
	commentModel := services.Comment{
		Id:            item.Id,
		User:          BuildUser(item.User),
		Content:		item.Content,
		CreateDate: 	item.CreatedAt.String(),
}
	return &commentModel
}


func BuildCommentList(item []models.Comment) []*services.Comment{
	commentlist := []*services.Comment{}
	for _,comment := range item{
		commentlist = append(commentlist,BuildComment(comment))
	}
	return commentlist
}

func (c CommentService) CommentAction(ctx context.Context, request *services.CommentActionRequest, response *services.CommentActionResponse) error {
	m, _ := jwt.ParseToken(request.Token)
	// 我的id
	myId := m.UserID

	actionType := request.ActionType

	if actionType==1{
		comment := &models.Comment{
			UserId: myId,
			VideoId: request.VideoId,
			Content: request.CommentText,
		}

		if err := mysql.AddComment(comment); err != nil {
			log.Println("AddComment failed")
			return nil
		}

		if err:= mysql.FindCommentUser(comment);err!=nil{
			log.Println("AddComment failed")
			return nil
		}

		// 对redis评论数zset中videoId ++
		redis.AddComment(comment)


		response.StatusCode = 0
		response.StatusMsg = "评论成功"

		response.Comment = BuildComment(*comment)


		return nil

	}else if actionType==2{

		videoId := request.VideoId
		comment := &models.Comment{
			Id: request.CommentId,
		}
		// 判断评论是否属于该用户
		if err:=mysql.CheckUser(comment,myId); err!=nil{
			response.StatusCode = 1
			response.StatusMsg = "用户无权限"
			return nil
		}

		// 删除
		if err := mysql.DelComment(comment); err != nil {
			response.StatusCode = 1
			response.StatusMsg = "删除失败"
			return nil
		}

		// 对redis评论数zset中videoId --
		redis.SubComment(videoId)


		response.StatusCode = 0
		response.StatusMsg = "删除成功"

		return nil
	}

	return nil
}

func (c CommentService) CommentList(ctx context.Context, request *services.CommentListRequest, response *services.CommentListResponse) error {

	videoId := request.VideoId

	CommentArr,err := mysql.GetCommentList(videoId)
	if err!=nil{
		response.StatusCode = 1
		response.StatusMsg = "获取评论列表失败"
		return nil
	}

	response.StatusCode=0;
	response.StatusMsg="获取评论列表成功"
	response.CommentList = BuildCommentList(CommentArr)

	return nil

}
