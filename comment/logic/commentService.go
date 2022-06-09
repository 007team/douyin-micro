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

func BuildComment(item models.Comment) *services.Comment {
	commentModel := services.Comment{
		Id:            item.Id,
		//User:          item.User,
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
		// 对redis评论数zset中videoId ++
		redis.AddComment(comment)

	}else if actionType==2{

	}




	return nil
}

func (c CommentService) CommentList(ctx context.Context, request *services.CommentListRequest, response *services.CommentListResponse) error {
	panic("implement me")
}
