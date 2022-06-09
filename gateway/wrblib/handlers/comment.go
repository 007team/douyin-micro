package handlers

import (
	"context"
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)
// 评论操作
func CommentAction(c *gin.Context){
	var CommentActionReq services.CommentActionRequest
	// 获取去params
	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Println("FollowList: videoId invalied", err)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "userinfo parseInt failed",
		})
		return
	}
	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.Atoi(actionTypeStr)
	if err != nil {
		log.Println("FollowList: actionType invalied", err)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "userinfo parseInt failed",
		})
		return
	}

	// 评论还是删除
	if actionType == 1{
		// 评论
		commentText := c.Query("commment_text")
		// 赋值
		CommentActionReq = services.CommentActionRequest{
			Token:      token,
			VideoId:    videoId,
			ActionType: int32(actionType),
			CommentText: commentText,

		}


	}else if actionType == 2{
		// 删除评论
		commentIdStr := c.Query("comment_id")
		commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
		if err != nil {
			log.Println("FollowList: videoId invalied", err)
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "userinfo parseInt failed",
			})
			return
		}
		// 赋值
		CommentActionReq = services.CommentActionRequest{
			Token: token,
			VideoId: videoId,
			ActionType: int32(actionType),
			CommentId: commentId,
		}

	}else{
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "系统错误",
		})
		return
	}

	// 从gin.Key中取出服务实例
	commentService := c.Keys["commentService"].(services.CommentService)
	CommentActionResp,err := commentService.CommentAction(context.Background(),&CommentActionReq)
	if err!=nil{
		log.Println("Login error:",err)
		return
	}
	c.JSON(http.StatusOK,CommentActionResp)


}