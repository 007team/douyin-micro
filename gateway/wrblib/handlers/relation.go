package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
)

// 关注列表
func FollowList(c *gin.Context) {
	var followReq services.FollowListRequest
	userIdStr := c.Query("user_id") // 获取用户id
	userid, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("FollowList: user_id invalied", err)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "userinfo parseInt failed",
		})
		return
	}
	followReq.UserId = userid
	// followReq.Token = c.Query("token")

	// 从gin.Key中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	followResp, err := userService.FollowList(context.Background(), &followReq)
	if err != nil {
		log.Println("Login error:", err)
		return
	}
	c.JSON(http.StatusOK, followResp)
}

//粉丝列表
func FollowerList(c *gin.Context) {
	var followerReq services.FollowerListRequest
	userIdStr := c.Query("user_id") // 获取用户id
	userid, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("FollowList: user_id invalied", err)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "FollowerList parseInt failed",
		})
		return
	}
	followerReq.UserId = userid

	// 从gin.Key中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	followerResp, err := userService.FollowerList(context.Background(), &followerReq)
	if err != nil {
		log.Println("Login error:", err)
		return
	}

	c.JSON(http.StatusOK, followerResp)

}
