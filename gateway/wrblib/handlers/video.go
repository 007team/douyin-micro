package handlers

import (
	"context"
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// feed流
func Feed(c *gin.Context) {
	var videoReq services.VideoFeedRequest
	videoReq.Token = c.Query("token")

	// 从gin.key 中取出服务实例
	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.Feed(context.Background(), &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)

}

// 发布视频
func Publish(c *gin.Context) {
	var videoReq services.VideoPublishActionRequest

	// 获取视频数据
	data, err := c.FormFile("data")
	if err != nil {
		log.Println("c.FormFile(\"data\") failed")
		return
	}

	//	将视频数据转换成 []dyte
	fileContent, _ := data.Open()
	var byteContainer []byte
	byteContainer = make([]byte, 1000000)
	fileContent.Read(byteContainer)

	// 参数写入 Request
	videoReq.Data = byteContainer
	videoReq.Title = c.PostForm("title")
	videoReq.Token = c.PostForm("token")

	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.PublishAction(c, &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)
}

// 用的发布视频列表
func PublishList(c *gin.Context) {
	var videoReq services.VideoPublishListRequest
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("strconv.ParseInt failed")
		return
	}
	videoReq.UserId = userId
	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.PublishList(context.Background(), &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)

}

// 点赞视频列表
func FavoriteList(c *gin.Context) {
	var videoReq services.VideoFavoriteListRequest
	userId, ok := c.Get("user_id")
	if !ok {
		log.Println("c.Get(\"user_id\") failed")
	}
	videoReq.UserId = userId.(int64)
	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.FavoriteList(context.Background(), &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)

}
